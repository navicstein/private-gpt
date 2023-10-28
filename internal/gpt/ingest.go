package gpt

import (
	"context"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"io"
	"navicstein/private-gpt/internal/database/model"
	"navicstein/private-gpt/internal/database/repository"
	"navicstein/private-gpt/internal/fileprocessing"
	"navicstein/private-gpt/internal/vectorstore"
	"os"
	"path/filepath"
	"strings"
)

func SplitDocument(document *model.Document) ([]schema.Document, error) {
	log.Debug().Msgf("splitting document: %+v", document.ID)

	splitter := textsplitter.NewRecursiveCharacter()
	if document.Meta == nil {
		document.Meta = make(map[string]any)
	}

	// always add the ids to metadata so that it can be filtered out later
	document.Meta["documentId"] = document.ID
	docs := []schema.Document{
		{
			PageContent: document.Text,
			Metadata:    document.Meta,
		},
	}

	splitted, err := textsplitter.SplitDocuments(splitter, docs)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("splitted document into %d chunks", len(splitted))

	if len(splitted) == 0 {
		return nil, errors.New("this document doesnt contain written text, size is 0")
	}

	return splitted, nil
}

// EmbedDocumentChunks embeds document chunks into vector store
func EmbedDocumentChunks(ctx context.Context, docChunks []schema.Document) error {
	log.Debug().Msgf("embedding %d chunks into vector store", len(docChunks))
	var (
		err error
	)

	vectorStore, err := vectorstore.GetVectorStore(ctx, vectorstore.GetDefaultCollectionName())
	if err != nil {
		return fmt.Errorf("failed to create vector store: %w", err)
	}

	log.Debug().Msgf("adding %d documents to vector store", len(docChunks))
	if err := vectorStore.AddDocuments(ctx, docChunks); err != nil {
		return fmt.Errorf("failed to add documents to vector store: %w", err)
	}

	log.Debug().Msg("successfully added documents to vector store")
	return nil
}

// ProcessDocumentFile takes a file in fs and extracts the text content out of it
func ProcessDocumentFile(ctx context.Context, filePath string) (*repository.CreateDocumentParams, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return nil, err
	}

	var (
		fStat, _    = file.Stat()
		fileExt     = strings.Replace(filepath.Ext(file.Name()), ".", "", 1)
		mimeType    = mime.String()
		fileContent string
	)

	switch {
	case mimeType == "application/pdf":
		fileContent, err = fileprocessing.ExtractTextFromPDF(ctx, file)

		// p := documentloaders.NewPDF(file, finfo.Size())
		// schemaDocs, err := p.Load(ctx)
		// if err != nil {
		//	return nil, fmt.Errorf("unable to load PDF file: %w", err)
		//}
		// for _, content := range schemaDocs {
		//	fileContent += content.PageContent
		//}
	case strings.Contains(mimeType, "video") || strings.Contains(mimeType, "audio"):
		fileContent, err = TranscribeMedia(ctx, filePath)
	case strings.Contains(mimeType, "text/html"):
		htmlLoader := documentloaders.NewHTML(file)
		schemaDocs, err := htmlLoader.Load(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, content := range schemaDocs {
			fileContent += content.PageContent
		}
	default:
		_, _ = file.Seek(0, io.SeekStart)
		fileContent, err = fileprocessing.GetTextFromFile(ctx, file)
	}
	if err != nil {
		return nil, err
	}

	name := filepath.Base(file.Name())
	docParams := repository.CreateDocumentParams{
		Name:     name,
		Type:     fileExt,
		MimeType: mimeType,
		Text:     fileContent,
		Size:     int(fStat.Size()),
	}
	return &docParams, nil
}
