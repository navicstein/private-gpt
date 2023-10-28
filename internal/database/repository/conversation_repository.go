package repository

import (
	"gorm.io/gorm"
	"navicstein/private-gpt/internal/database/model"
)

type FindConversationsParams struct {
	Skip  int `json:"skip,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type ConversationRepositoryInterface interface {
	FindConversations(FindConversationsParams) ([]*model.Conversation, error)
	CreateMessage(role, message string) (*model.Conversation, error)
}

type Conversation struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepositoryInterface {
	return &Conversation{
		db: db,
	}
}

func (c Conversation) FindConversations(params FindConversationsParams) ([]*model.Conversation, error) {
	var (
		messages []*model.Conversation
	)

	query := c.db.Model(&model.Conversation{})

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Skip > 0 {
		query = query.Offset(params.Skip)
	}

	query = query.Order("created_at asc")

	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (c Conversation) CreateMessage(role, text string) (*model.Conversation, error) {
	var (
		message = model.Conversation{
			Text: text,
			Role: role,
		}
	)
	if err := c.db.Create(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}
