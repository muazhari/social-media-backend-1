package container

import "social-media-backend-1/internal/inners/use_cases"

type UseCaseContainer struct {
	AccountUseCase *use_cases.AccountUseCase
}

func NewUseCaseContainer(repositoryContainer *RepositoryContainer) *UseCaseContainer {
	return &UseCaseContainer{
		AccountUseCase: use_cases.NewAccountUseCase(repositoryContainer.AccountRepository),
	}
}
