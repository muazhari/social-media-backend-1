package entities

import "github.com/google/uuid"

type Account struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string
	TotalPostLike    int `json:"total_post_like"`
	TotalChatMessage int `json:"total_chat_message"`
}
