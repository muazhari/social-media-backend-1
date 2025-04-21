package entities

import "github.com/google/uuid"

type PostLike struct {
	ID        *uuid.UUID `json:"id"`
	AccountID uuid.UUID  `json:"account_id"`
	Post      *Post      `json:"post"`
}
