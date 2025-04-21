package repositories

import (
	"encoding/json"
	"fmt"
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
	query := `
		SELECT COALESCE(json_agg(json_build_object(
			'id', id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message
		)), '[]'::json) as item
		FROM account;
	`

	var jsonData []byte

	err := r.TwoDatabaseConfig.Connection.QueryRow(query).Scan(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("database query scan failed: %w", err)
	}

	var accounts []*entities.Account
	err = json.Unmarshal(jsonData, &accounts)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling accounts JSON: %w", err)
	}

	return accounts, nil
}

func (r *AccountRepository) GetAccountById(id uuid.UUID) (*entities.Account, error) {
	query := `
		SELECT json_build_object(
			'id', id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message
		) as item
		FROM account
		WHERE account.id = $1
	`

	var jsonData []byte
	err := r.TwoDatabaseConfig.Connection.QueryRow(query, id).Scan(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("database query scan failed: %w", err)
	}

	var account *entities.Account
	err = json.Unmarshal(jsonData, &account)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling account JSON: %w", err)
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	return account, nil
}

func (r *AccountRepository) GetAccountsByIds(ids []*uuid.UUID) ([]*entities.Account, []error) {
	query := `
		SELECT COALESCE(json_agg(json_build_object(
			'id', id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message
		)), '[]'::json) as item
		FROM account
		WHERE id IN ($1)
	`

	var jsonData []byte
	err := r.TwoDatabaseConfig.Connection.QueryRow(query, ids).Scan(&jsonData)
	if err != nil {
		return nil, []error{fmt.Errorf("database query scan failed: %w", err)}
	}

	var accounts []*entities.Account
	err = json.Unmarshal(jsonData, &accounts)
	if err != nil {
		return nil, []error{fmt.Errorf("error unmarshaling accounts JSON: %w", err)}
	}

	if len(accounts) == 0 {
		return nil, []error{fmt.Errorf("accounts not found")}
	}

	return accounts, nil
}

func (r *AccountRepository) CreateAccount(account *entities.Account) (*entities.Account, error) {
	query := `
		INSERT INTO account (id, name, email, password, total_post_like, total_chat_message)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.TwoDatabaseConfig.Connection.Exec(
		query,
		account.ID,
		account.Name,
		account.Email,
		account.Password,
		account.TotalPostLike,
		account.TotalChatMessage,
	)
	if err != nil {
		return nil, fmt.Errorf("database insert failed: %w", err)
	}

	createdAccount, err := r.GetAccountById(*account.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created account: %w", err)
	}

	return createdAccount, nil
}

func (r *AccountRepository) UpdateAccountById(id uuid.UUID, account *entities.Account) (*entities.Account, error) {
	query := `
		UPDATE account
		SET name = $1, email = $2, password = $3, total_post_like = $4, total_chat_message = $5
		WHERE id = $6
	`

	_, err := r.TwoDatabaseConfig.Connection.Exec(
		query,
		account.Name,
		account.Email,
		account.Password,
		account.TotalPostLike,
		account.TotalChatMessage,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("database update failed: %w", err)
	}

	updatedAccount, err := r.GetAccountById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated account: %w", err)
	}

	return updatedAccount, nil
}

func (r *AccountRepository) DeleteAccountById(id uuid.UUID) (*entities.Account, error) {
	query := `
		DELETE FROM account
		WHERE id = $1
	`

	_, err := r.TwoDatabaseConfig.Connection.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("database delete failed: %w", err)
	}

	return nil, nil
}
