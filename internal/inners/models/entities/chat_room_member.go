package entities

import "github.com/google/uuid"

type ChatRoomMember struct {
	ID        *uuid.UUID `json:"id"`
	AccountID uuid.UUID  `json:"account_id"`
	ChatRoom  *ChatRoom  `json:"chat_room"`
}
