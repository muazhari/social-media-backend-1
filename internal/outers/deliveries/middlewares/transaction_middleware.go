package middlewares

import (
	"social-media-backend-1/internal/outers/repositories"
)

type TransactionMiddleware struct {
}

func NewTransactionMiddleware(
	accountRepository *repositories.AccountRepository,
) *TransactionMiddleware {
	return &TransactionMiddleware{}
}
