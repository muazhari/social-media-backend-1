package container

import "social-media-backend-1/internal/inners/use_cases"

type UseCaseContainer struct {
	AccountUseCase *use_cases.AccountUseCase
	AuthUseCase    *use_cases.AuthUseCase
}

func NewUseCaseContainer(repositoryContainer *RepositoryContainer, gatewayContainer *GatewayContainer) *UseCaseContainer {
	return &UseCaseContainer{
		AccountUseCase: use_cases.NewAccountUseCase(repositoryContainer.AccountRepository, repositoryContainer.FileRepository),
		AuthUseCase:    use_cases.NewAuthUseCase(repositoryContainer.AccountRepository, gatewayContainer.AuthGateway),
	}
}
