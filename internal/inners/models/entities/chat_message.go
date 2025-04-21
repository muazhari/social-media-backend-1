package entities

import "github.com/google/uuid"

type ChatMessage struct {
	ID        *uuid.UUID `json:"id"`
	Content   string     `json:"content"`
	AccountID uuid.UUID  `json:"account_id"`
	ChatRoom  *ChatRoom  `json:"chat_room"`
}
