package container

type RootContainer struct {
	ConfigContainer     *ConfigContainer
	RepositoryContainer *RepositoryContainer
	UseCaseContainer    *UseCaseContainer
	MiddlewareContainer *MiddlewareContainer
}

func NewRootContainer() *RootContainer {
	configContainer := NewConfigContainer()
	repositoryContainer := NewRepositoryContainer(configContainer)
	useCaseContainer := NewUseCaseContainer(repositoryContainer)
	middlewareContainer := NewMiddlewareContainer(repositoryContainer)

	return &RootContainer{
		ConfigContainer:     configContainer,
		RepositoryContainer: repositoryContainer,
		UseCaseContainer:    useCaseContainer,
		MiddlewareContainer: middlewareContainer,
	}
}
