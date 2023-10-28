package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/vectorstores"
	"navicstein/private-gpt/internal/config"
	"navicstein/private-gpt/internal/helpers"
	"navicstein/private-gpt/internal/vectorstore"
	"strings"
)

type QueryIngestParams struct {
	CollectionIDs []string `json:"collectionIds"`
	Query         any      `json:"query"`
}

type ChatParams struct {
	NewMessage string
	DocumentID string // if chatting in the scope of a single document
	//CollectionID string // if chatting in the scope of collections
}

type ChatToken string // the token to be used for chatting

const (
	StartToken    = "START"
	EndToken      = "END"
	ErrNoResponse = "There's a problem, i can't help with that"
)

// SocketMessageResponse the response from the socket
type SocketMessageResponse struct {
	IsBuffered bool      `json:"isBuffered"`
	Token      ChatToken `json:"token,omitempty"`
	Sources    any       `json:"sources,omitempty"`
	Message    string    `json:"message,omitempty"`
}

// EmitMessage sends the message to the client
func EmitMessage(channel chan []byte, message SocketMessageResponse) {
	channel <- helpers.MustMarshalJSON(message)
}

// ChatWithStore chat's with a collection or a document for a user
// ChatWithStore is a query + chat history with streaming output
func (s *Service) ChatWithStore(ctx context.Context, params ChatParams, messageCh chan []byte) {
	var (
		responseBuffer bytes.Buffer
		cfg            = config.GetConfig()
	)

	// handleResponseStream sends the response to the client
	var handleResponseStream = func(ctx context.Context, chunk []byte) error { // accumulate all chunks and send as a whole
		if len(chunk) == 0 {
			return nil
		}

		_, err := responseBuffer.Write(chunk)
		if err != nil {
			return fmt.Errorf("failed to write chunk to buffer: %w", err)
		}

		// format the response in markdown as HTML
		htmlMessage := helpers.MDToHTML(responseBuffer.String())
		responseStream := SocketMessageResponse{
			IsBuffered: true,
			Message:    htmlMessage,
		}
		EmitMessage(messageCh, responseStream)
		return nil
	}

	llm, err := openai.NewChat(openai.WithBaseURL(cfg.OpenAI.BaseURL), openai.WithToken(cfg.OpenAI.APIKey))
	if err != nil {
		log.Err(err).Msg("failed to create llm")
		EmitMessage(messageCh, SocketMessageResponse{
			Message: ErrNoResponse,
		})
		return
	}

	var (
		numOfDocuments = 2
	)

	// 1. load the related documents from the vectorStore
	vectorStore, err := vectorstore.GetVectorStore(ctx, vectorstore.GetDefaultCollectionName())
	if err != nil {
		log.Err(err).Msg("failed to get vector store")
		return
	}

	docs, err := vectorStore.SimilaritySearch(ctx, params.NewMessage, numOfDocuments,
		vectorstores.WithScoreThreshold(0.7),
	)

	if err != nil {
		if strings.Contains(err.Error(), "empty response") {
			// send msg to the chat client
			EmitMessage(messageCh, SocketMessageResponse{
				Message: ErrNoResponse,
			})
			log.Err(err).Msg("[vectorStore] empty response")
			return
		}

		EmitMessage(messageCh, SocketMessageResponse{
			Message: ErrNoResponse,
		})
		log.Err(err).Msg("[vectorStore] failed to search for documents")
		return
	}

	if err != nil {
		log.Err(err).Msg("failed to find chat history")
		return
	}

	_conversationTemplate := `Answer the following question from the document. If you don't not know the answer to a question, simply apologize then try to come up with an answer from the document, Your answer must be in a markdown format

	## document:
	what can be done here
	
	question: {{.question}}
	`
	//{{.context}}

	//TODO: use https://github.com/navicstein/langchaingo/blob/main/chains/conversational_retrieval_qa_test.go
	customPrompt := prompts.NewPromptTemplate(_conversationTemplate, []string{"question", "context"})
	promptMap := map[string]any{
		"input_documents": docs,
		"question":        params.NewMessage,
	}

	llmChain := chains.NewLLMChain(llm, customPrompt)
	chain := chains.NewStuffDocuments(llmChain)

	_, err = s.conversationRepo.CreateMessage("human", params.NewMessage)
	if err != nil {
		log.Err(err).Msg("failed to create chat human message")
	}

	EmitMessage(messageCh, SocketMessageResponse{
		IsBuffered: true,
		Token:      StartToken,
	})
	result, err := chains.Call(ctx, chain, promptMap, chains.WithStreamingFunc(handleResponseStream))
	if err != nil {
		log.Err(err).Msg("failed to run chains")

		//TODO: This is a hack to get the chat type out the error message
		for _, msg := range strings.Fields(err.Error()) {
			_ = handleResponseStream(ctx, []byte(msg+"\n"))
		}
		EmitMessage(messageCh, SocketMessageResponse{
			IsBuffered: true,
			Token:      EndToken,
		})
		return
	}

	aiMsg, ok := result["text"].(string)
	if !ok {
		panic("llm resp is not a string")
	}

	_, err = s.conversationRepo.CreateMessage("ai", aiMsg)
	if err != nil {
		log.Err(err).Msg("failed to create chat human message")
	}

	EmitMessage(messageCh, SocketMessageResponse{
		IsBuffered: true,
		Token:      EndToken,
	})

	log.Info().Msg(aiMsg)
}

// QueryStoreForText custom prompt for text output
func (s *Service) QueryStoreForText(ctx context.Context, query QueryIngestParams) (string, error) {
	cfg := config.GetConfig()

	vectorStore, err := vectorstore.GetVectorStore(ctx, vectorstore.GetDefaultCollectionName())
	if err != nil {
		return "", fmt.Errorf("failed to create vector store: %w", err)
	}

	llm, err := openai.New(openai.WithBaseURL(cfg.OpenAI.BaseURL), openai.WithToken(cfg.OpenAI.APIKey))
	if err != nil {
		return "", err
	}

	vectorRetriever := vectorstores.ToRetriever(vectorStore, 2,
		vectorstores.WithScoreThreshold(1),
	)

	result, err := chains.Run(ctx,
		chains.NewRetrievalQAFromLLM(llm, vectorRetriever),
		query.Query,
	)
	if err != nil {
		return "", fmt.Errorf("failed to run query chains: %w", err)
	}

	return result, nil
}
