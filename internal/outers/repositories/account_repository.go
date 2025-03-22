package repositories

import (
	"social-media-backend-1/internal/outers/configs"
)

type AccountRepository struct {
	TwoDatabaseConfig *configs.TwoDatabaseConfig
}

func NewAccountRepository(twoDatabaseConfig *configs.TwoDatabaseConfig) *AccountRepository {
	return &AccountRepository{
		TwoDatabaseConfig: twoDatabaseConfig,
	}
}
