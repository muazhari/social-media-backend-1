package container

import (
	"social-media-backend-1/internal/outers/deliveries/middlewares"
)

type MiddlewareContainer struct {
	AuthMiddleware        *middlewares.AuthMiddleware
	TransactionMiddleware *middlewares.TransactionMiddleware
}

func NewMiddlewareContainer(
	useCaseContainer *UseCaseContainer,
	repositoryContainer *RepositoryContainer,
) *MiddlewareContainer {
	return &MiddlewareContainer{
		AuthMiddleware:        middlewares.NewAuthMiddleware(useCaseContainer.AuthUseCase),
		TransactionMiddleware: middlewares.NewTransactionMiddleware(repositoryContainer.AccountRepository),
	}
}
