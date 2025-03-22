package container

import (
	"social-media-backend-1/internal/outers/repositories"
)

type RepositoryContainer struct {
	AccountRepository *repositories.AccountRepository
}

func NewRepositoryContainer(configContainer *ConfigContainer) *RepositoryContainer {
	return &RepositoryContainer{
		AccountRepository: repositories.NewAccountRepository(configContainer.TwoDatabaseConfig),
	}
}
