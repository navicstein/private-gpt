package repository

import (
	"errors"
	"gorm.io/gorm"
	"navicstein/private-gpt/internal/database/model"
	"strings"
)

type CreateDocumentParams struct {
	Name          string `json:"name"`
	Type          string
	MimeType      string
	Text          string
	Size          int
	Src           string
	HasTranscript bool
	Meta          map[string]any
}

type FindDocumentParams struct {
	Skip  int `json:"skip,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type DocumentRepositoryInterface interface {
	FindDocuments(params FindDocumentParams) ([]*model.Document, error)
	FindDocumentByID(id string) (*model.Document, error)
	UpdateDocument(document *model.Document) error
	DeleteDocument(document *model.Document) error
	CreateDocument(document CreateDocumentParams) (*model.Document, error)
}

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepositoryInterface {
	return &DocumentRepository{
		db: db,
	}
}

func (r *DocumentRepository) CreateDocument(params CreateDocumentParams) (*model.Document, error) {
	var (
		newDocument = model.Document{
			Name:     strings.ToLower(params.Name),
			Size:     params.Size,
			Type:     params.Type,
			MimeType: params.MimeType,
			Text:     params.Text,
			Meta:     params.Meta,
		}
	)
	if err := r.db.Create(&newDocument).Error; err != nil {
		return nil, err
	}

	return &newDocument, nil
}

func (r *DocumentRepository) UpdateDocument(document *model.Document) error {
	if document.ID == "" {
		return errors.New("document must not be empty")
	}
	return r.db.Model(&document).Updates(&document).Error
}

func (r *DocumentRepository) DeleteDocument(document *model.Document) error {
	return r.db.Model(model.Document{}).Unscoped().Delete(document).Error
}

// FindDocuments finds all documents for a collection or a user with pagination
func (r *DocumentRepository) FindDocuments(params FindDocumentParams) ([]*model.Document, error) {
	var (
		documents []*model.Document
	)

	query := r.db.Model(&model.Document{})

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Skip > 0 {
		query = query.Offset(params.Skip)
	}

	query = query.Order("created_at desc")

	if err := query.Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (r *DocumentRepository) FindDocumentByID(id string) (*model.Document, error) {
	var (
		document model.Document
	)
	if err := r.db.Model(&model.Document{}).Where("id = ?", id).Preload("User").First(&document).Error; err != nil {
		return nil, err
	}
	return &document, nil
}
