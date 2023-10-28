package model

type Document struct {
	Model
	Name     string         `json:"name,omitempty"`
	Size     int            `json:"size,omitempty"`
	Type     string         `json:"type,omitempty"`
	MimeType string         `json:"mimeType,omitempty"`
	Text     string         `json:"text,omitempty"`
	Meta     map[string]any `gorm:"serializer:json" json:"meta,omitempty"`
}
