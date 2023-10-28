package model

// Conversation TODO: is a conversation between two users
type Conversation struct {
	Model
	Text string `json:"text,omitempty"`
	Role string `json:"role,omitempty"`
}
