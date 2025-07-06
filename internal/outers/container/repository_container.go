package container

import (
	"social-media-backend-1/internal/outers/repositories"
)

type RepositoryContainer struct {
	AccountRepository *repositories.AccountRepository
	FileRepository    *repositories.FileRepository
}

func NewRepositoryContainer(configContainer *ConfigContainer) *RepositoryContainer {
	return &RepositoryContainer{
		AccountRepository: repositories.NewAccountRepository(configContainer.TwoDatastoreConfig),
		FileRepository:    repositories.NewFileRepository(configContainer.FourDatastoreConfig),
	}
}
