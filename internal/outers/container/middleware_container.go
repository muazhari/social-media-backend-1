package container

import (
	"social-media-backend-1/internal/outers/deliveries/middlewares"
)

type MiddlewareContainer struct {
	TransactionMiddleware *middlewares.TransactionMiddleware
}

func NewMiddlewareContainer(repositoryContainer *RepositoryContainer) *MiddlewareContainer {
	return &MiddlewareContainer{
		TransactionMiddleware: middlewares.NewTransactionMiddleware(repositoryContainer.AccountRepository),
	}
}
