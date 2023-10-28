package service

import (
	"gorm.io/gorm"
	"navicstein/private-gpt/internal/database/repository"
)

type Service struct {
	db               *gorm.DB
	documentRepo     repository.DocumentRepositoryInterface
	conversationRepo repository.ConversationRepositoryInterface
}

// NewService creates a new service
// TODO: add repository dependency
func NewService(db *gorm.DB) Service {
	return Service{
		conversationRepo: repository.NewConversationRepository(db),
		documentRepo:     repository.NewDocumentRepository(db),
	}
}
