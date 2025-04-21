package entities

import "github.com/google/uuid"

type ChatRoom struct {
	ID              *uuid.UUID        `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	ChatMessages    []*ChatMessage    `json:"chat_message"`
	ChatRoomMembers []*ChatRoomMember `json:"chat_room_member"`
}
