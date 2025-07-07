package value_objects

import (
	"social-media-backend-1/internal/inners/models/entities"
)

type Session struct {
	Account      *entities.Account
	AccessToken  string
	RefreshToken string
}
