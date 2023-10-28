package api

import (
	"context"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"navicstein/private-gpt/internal/api/service"
	"navicstein/private-gpt/internal/database/repository"
	"navicstein/private-gpt/internal/helpers"
	"net/http"
)

func (a *API) QueryData(ctx *fiber.Ctx) error {
	var (
		params service.QueryIngestParams
	)
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	response, err := a.service.QueryStoreForText(ctx.Context(), params)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.JSON(response)
}

func (a *API) FindDocuments(ctx *fiber.Ctx) error {
	// TODO: define from client
	params := repository.FindDocumentParams{
		Skip:  0,
		Limit: 100,
	}

	documents, err := a.documentRepository.FindDocuments(params)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.JSON(documents)
}

func (a *API) FindConversations(ctx *fiber.Ctx) error {
	// TODO: define from client
	params := repository.FindConversationsParams{
		Skip:  0,
		Limit: 100,
	}

	messages, err := a.conversationRepository.FindConversations(params)

	for _, msg := range messages {
		if msg.Role == "ai" {
			msg.Text = helpers.MDToHTML(msg.Text)
		}
	}

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.JSON(messages)
}

// WebsocketChatHandler is the controller for chatting with an entity via websockets
func (a *API) WebsocketChatHandler(c *websocket.Conn) {
	//TODO: clean this shitty line
	ctx, cancel := context.WithCancel(context.Background())

	if allowed, ok := c.Locals("allowed").(bool); !ok || !allowed {
		log.Error().Msg("connection needs to be upgraded!")
		return
	}

	var (
		mt  = websocket.TextMessage
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Warn().Msg("error reading message from websocket:" + err.Error())
			}
			cancel()
			_ = c.Close()
			break
		}

		params := service.ChatParams{
			NewMessage: string(msg),
		}

		replyCh := make(chan []byte)

		go func() {
			a.service.ChatWithStore(ctx, params, replyCh)
			close(replyCh)
		}()

		go func() {
			for message := range replyCh {
				if err = c.WriteMessage(mt, message); err != nil {
					log.Err(err).Msg("error writing message to websocket")
					cancel()
					break
				}
			}
		}()
	}
}
