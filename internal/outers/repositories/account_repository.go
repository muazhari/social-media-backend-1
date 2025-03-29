package repositories

import (
	"github.com/google/uuid"
	"social-media-backend-1/internal/inners/models/entities"
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

func (r *AccountRepository) GetAllAccounts() ([]*entities.Account, error) {
	return r.TwoDatabaseConfig.AccountData, nil
}

func (r *AccountRepository) GetAccountById(id uuid.UUID) (*entities.Account, error) {
	for _, account := range r.TwoDatabaseConfig.AccountData {
		if account.ID == id {
			return account, nil
		}
	}
	return nil, nil
}

func (r *AccountRepository) CreateAccount(account *entities.Account) error {
	r.TwoDatabaseConfig.AccountData = append(r.TwoDatabaseConfig.AccountData, account)
	return nil
}

func (r *AccountRepository) UpdateAccountById(id uuid.UUID, account *entities.Account) (*entities.Account, error) {
	for i, oldAccount := range r.TwoDatabaseConfig.AccountData {
		if oldAccount.ID == id {
			r.TwoDatabaseConfig.AccountData[i] = account
			return account, nil
		}
	}
	return nil, nil
}

func (r *AccountRepository) DeleteAccountById(id uuid.UUID) (*entities.Account, error) {
	for i, account := range r.TwoDatabaseConfig.AccountData {
		if account.ID == id {
			r.TwoDatabaseConfig.AccountData = append(r.TwoDatabaseConfig.AccountData[:i], r.TwoDatabaseConfig.AccountData[i+1:]...)
			return account, nil
		}
	}
	return nil, nil
}
