package entities

import "github.com/google/uuid"

type Account struct {
	ID               *uuid.UUID     `json:"id"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	TotalPostLike    float64        `json:"total_post_like"`
	TotalChatMessage float64        `json:"total_chat_message"`
	Scopes           []string       `json:"scopes"`
	Messages         []*ChatMessage `json:"messages"`
	Rooms            []*ChatRoom    `json:"rooms"`
	Posts            []*Post        `json:"posts"`
	PostLikes        []*PostLike    `json:"post_likes"`
}
