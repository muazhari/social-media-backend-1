package use_cases

import (
	"github.com/google/uuid"
	"social-media-backend-1/internal/inners/models/entities"
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

func (uc *AccountUseCase) GetAllAccounts() ([]*entities.Account, error) {
	accounts, err := uc.AccountRepository.GetAllAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (uc *AccountUseCase) GetAccountById(id uuid.UUID) (*entities.Account, error) {
	account, err := uc.AccountRepository.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (uc *AccountUseCase) CreateAccount(account *entities.Account) (*entities.Account, error) {
	account, err := uc.AccountRepository.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (uc *AccountUseCase) UpdateAccountById(id uuid.UUID, account *entities.Account) (*entities.Account, error) {
	account, err := uc.AccountRepository.UpdateAccountById(id, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (uc *AccountUseCase) DeleteAccountById(id uuid.UUID) (*entities.Account, error) {
	account, err := uc.AccountRepository.DeleteAccountById(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
