package entities

import "github.com/google/uuid"

type Post struct {
	ID        *uuid.UUID  `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	AccountID uuid.UUID   `json:"account_id"`
	PostLikes []*PostLike `json:"post_likes"`
}
