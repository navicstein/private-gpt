package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"navicstein/private-gpt/internal/api/service"
	"navicstein/private-gpt/internal/database/repository"
	"time"
)

type API struct {
	app                    *fiber.App
	service                service.Service
	documentRepository     repository.DocumentRepositoryInterface
	conversationRepository repository.ConversationRepositoryInterface
	db                     *gorm.DB
}

func NewAPI(app *fiber.App, db *gorm.DB) API {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "X-Requested-With, Origin, Cache-Control, Content-Type, Accept, Access-Control-Allow-Headers, Authorization",
	}))

	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 42 * time.Minute,
	}))

	app.Use(helmet.New())
	app.Use(favicon.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	return API{
		app: app,

		service:                service.NewService(db),
		documentRepository:     repository.NewDocumentRepository(db),
		conversationRepository: repository.NewConversationRepository(db),
		db:                     db,
	}
}

// RegisterRoutes
// ╦═╗┌─┐┬ ┬┌┬┐┌─┐┌─┐┬
// ╠╦╝│ ││ │ │ ├┤ └─┐│
// ╩╚═└─┘└─┘ ┴ └─┘└─┘o
func (a *API) RegisterRoutes() {
	var (
		app = a.app
	)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "privateGPT Metrics"}))

	v1 := app.Group("/v1")
	v1.Use("/chat", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		log.Error().Msg("connection needs to be upgraded!")
		return fiber.ErrUpgradeRequired
	})

	v1.Get("/chat", websocket.New(a.WebsocketChatHandler))
	v1.Post("/ingest", a.IngestData)
	v1.Get("/query", a.QueryData)

	v1.Get("/documents", a.FindDocuments)
	v1.Get("/conversations", a.FindConversations)
}
