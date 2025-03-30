package entities

import "github.com/google/uuid"

type Account struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string
	TotalPostLike    float64 `json:"total_post_like"`
	TotalChatMessage float64 `json:"total_chat_message"`
}
