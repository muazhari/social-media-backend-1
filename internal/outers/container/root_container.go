package container

type RootContainer struct {
	ConfigContainer     *ConfigContainer
	RepositoryContainer *RepositoryContainer
	UseCaseContainer    *UseCaseContainer
}

func NewRootContainer() *RootContainer {
	configContainer := NewConfigContainer()
	repositoryContainer := NewRepositoryContainer(configContainer)
	useCaseContainer := NewUseCaseContainer(repositoryContainer)

	return &RootContainer{
		ConfigContainer:     configContainer,
		RepositoryContainer: repositoryContainer,
		UseCaseContainer:    useCaseContainer,
	}
}
