package use_cases

import (
	"social-media-backend-1/internal/outers/repositories"
)

type AccountUseCase struct {
	AccountRepository *repositories.AccountRepository
}

func NewAccountUseCase(accountRepository *repositories.AccountRepository) *AccountUseCase {
	return &AccountUseCase{
		AccountRepository: accountRepository,
	}
}
