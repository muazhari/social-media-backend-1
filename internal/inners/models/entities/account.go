package entities

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type Account struct {
	ID               *uuid.UUID      `json:"id"`
	ImageID          *uuid.UUID      `json:"image_id"`
	ImageURL         *string         `json:"image_url"`
	Image            *graphql.Upload `json:"image"`
	Name             *string         `json:"name"`
	Email            *string         `json:"email"`
	Password         *string         `json:"password"`
	TotalPostLike    *float64        `json:"total_post_like"`
	TotalChatMessage *float64        `json:"total_chat_message"`
	Scopes           []string        `json:"scopes"`
	Messages         []*ChatMessage  `json:"messages"`
	Rooms            []*ChatRoom     `json:"rooms"`
	Posts            []*Post         `json:"posts"`
	PostLikes        []*PostLike     `json:"post_likes"`
}
