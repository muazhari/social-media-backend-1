package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"social-media-backend-1/internal/inners/models/entities"
	"social-media-backend-1/internal/outers/configs"

	"github.com/google/uuid"
)

type AccountRepository struct {
	TwoDatastoreConfig *configs.TwoDatastoreConfig
}

func NewAccountRepository(twoDatabaseConfig *configs.TwoDatastoreConfig) *AccountRepository {
	return &AccountRepository{
		TwoDatastoreConfig: twoDatabaseConfig,
	}
}

func (r *AccountRepository) GetAllAccounts(ctx context.Context) ([]*entities.Account, error) {
	query := `
		SELECT COALESCE(json_agg(json_build_object(
			'id', id,
			'image_id', image_id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message,
		    'scopes', (
				SELECT COALESCE(json_agg(account_scope.scope) , '[]'::json) 
				FROM account_scope 
				WHERE account_scope.account_id = account.id
		    ) 
		)), '[]'::json)
		FROM account;
	`

	var jsonData []byte
	err := r.TwoDatastoreConfig.Connection.QueryRowContext(ctx, query).Scan(&jsonData)
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

func (r *AccountRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	query := `
		SELECT json_build_object(
			'id', id,
			'image_id', image_id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message,
		    'scopes', (
				SELECT COALESCE(json_agg(account_scope.scope) , '[]'::json) 
				FROM account_scope 
				WHERE account_scope.account_id = account.id
		    )
		)
		FROM account
		WHERE account.id = $1
	`

	var jsonData []byte
	err := r.TwoDatastoreConfig.Connection.QueryRowContext(ctx, query, id).Scan(&jsonData)
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

func (r *AccountRepository) GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*entities.Account, error) {
	query := `
		SELECT json_build_object(
			'id', id,
			'image_id', image_id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message,
		    'scopes', (
				SELECT COALESCE(json_agg(account_scope.scope) , '[]'::json) 
				FROM account_scope 
				WHERE account_scope.account_id = account.id
		    )
		)
		FROM account
		WHERE email = $1 AND password = $2
	`

	var jsonData []byte
	err := r.TwoDatastoreConfig.Connection.QueryRowContext(ctx, query, email, password).Scan(&jsonData)
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

func (r *AccountRepository) GetAccountsByIDs(ctx context.Context, ids []*uuid.UUID) ([]*entities.Account, error) {
	query := `
		SELECT COALESCE(json_agg(json_build_object(
			'id', id,
			'image_id', image_id,
			'name', name,
			'email', email,
			'password', password,
			'total_post_like', total_post_like,
			'total_chat_message', total_chat_message,
		    'scopes', (
				SELECT COALESCE(json_agg(account_scope.scope) , '[]'::json) 
				FROM account_scope 
				WHERE account_scope.account_id = account.id
		    )
		)), '[]'::json)
		FROM account
		WHERE id = ANY($1)
	`

	var jsonData []byte
	err := r.TwoDatastoreConfig.Connection.QueryRowContext(ctx, query, ids).Scan(&jsonData)
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

func (r *AccountRepository) CreateAccount(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	query_one := `
		INSERT INTO account (id, image_id, name, email, password, total_post_like, total_chat_message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.TwoDatastoreConfig.Connection.ExecContext(
		ctx,
		query_one,
		account.ID,
		account.ImageID,
		account.Name,
		account.Email,
		account.Password,
		account.TotalPostLike,
		account.TotalChatMessage,
	)
	if err != nil {
		return nil, fmt.Errorf("database insert one failed: %w", err)
	}

	queryTwo := `
			INSERT INTO account_scope (account_id, scope)
			SELECT $1, unnest($2::text[])
		`

	_, err = r.TwoDatastoreConfig.Connection.ExecContext(
		ctx,
		queryTwo,
		account.ID,
		account.Scopes,
	)
	if err != nil {
		return nil, fmt.Errorf("database insert two failed: %w", err)
	}

	createdAccount, err := r.GetAccountByID(ctx, *account.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created account: %w", err)
	}

	return createdAccount, nil
}

func (r *AccountRepository) UpdateAccountByID(ctx context.Context, id uuid.UUID, account *entities.Account) (*entities.Account, error) {
	query_one := `
		UPDATE account
		SET image_id = $1, name = $2, email = $3, password = $4, total_post_like = $5, total_chat_message = $6
		WHERE id = $7
	`

	_, err := r.TwoDatastoreConfig.Connection.ExecContext(
		ctx,
		query_one,
		account.ImageID,
		account.Name,
		account.Email,
		account.Password,
		account.TotalPostLike,
		account.TotalChatMessage,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("database update one failed: %w", err)
	}

	queryTwo := `
		DELETE FROM account_scope
		WHERE account_id = $1
	`

	_, err = r.TwoDatastoreConfig.Connection.ExecContext(ctx, queryTwo, id)
	if err != nil {
		return nil, fmt.Errorf("database delete two failed: %w", err)
	}

	queryThree := `
		INSERT INTO account_scope (account_id, scope)
		SELECT $1, unnest($2::text[])
	`

	_, err = r.TwoDatastoreConfig.Connection.ExecContext(
		ctx,
		queryThree,
		id,
		account.Scopes,
	)
	if err != nil {
		return nil, fmt.Errorf("database insert three failed: %w", err)
	}

	updatedAccount, err := r.GetAccountByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated account: %w", err)
	}

	return updatedAccount, nil
}

func (r *AccountRepository) DeleteAccountByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	foundAccount, err := r.GetAccountByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve account for deletion: %w", err)
	}

	query := `
		DELETE FROM account
		WHERE id = $1
	`

	_, err = r.TwoDatastoreConfig.Connection.ExecContext(ctx, query, foundAccount.ID)
	if err != nil {
		return nil, fmt.Errorf("database delete failed: %w", err)
	}

	return foundAccount, nil
}

func (r *AccountRepository) IncrementTotalPostLike(ctx context.Context, accountID uuid.UUID, delta float64) error {
	query := `
		UPDATE account
		SET total_post_like = COALESCE(total_post_like, 0) + $1
		WHERE id = $2::uuid
	`
	_, err := r.TwoDatastoreConfig.Connection.ExecContext(ctx, query, delta, accountID)
	if err != nil {
		return fmt.Errorf("database update total_post_like failed: %w", err)
	}
	return nil
}

func (r *AccountRepository) DecrementTotalPostLike(ctx context.Context, accountID uuid.UUID, delta float64) error {
	query := `
		UPDATE account
		SET total_post_like = COALESCE(total_post_like, 0) - $1
		WHERE id = $2::uuid
	`
	_, err := r.TwoDatastoreConfig.Connection.ExecContext(ctx, query, delta, accountID)
	if err != nil {
		return fmt.Errorf("database update total_post_like failed: %w", err)
	}
	return nil
}

func (r *AccountRepository) IncrementTotalChatMessage(ctx context.Context, accountID uuid.UUID, delta float64) error {
	query := `
		UPDATE account
		SET total_chat_message = COALESCE(total_chat_message, 0) + $1
		WHERE id = $2::uuid
	`
	_, err := r.TwoDatastoreConfig.Connection.ExecContext(ctx, query, delta, accountID)
	if err != nil {
		return fmt.Errorf("database update total_chat_message failed: %w", err)
	}
	return nil
}
