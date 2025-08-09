package container

type RootContainer struct {
	ConfigContainer     *ConfigContainer
	RepositoryContainer *RepositoryContainer
	GatewayContainer    *GatewayContainer
	UseCaseContainer    *UseCaseContainer
	MiddlewareContainer *MiddlewareContainer
}

func NewRootContainer() *RootContainer {
	configContainer := NewConfigContainer()
	repositoryContainer := NewRepositoryContainer(configContainer)
	gatewayContainer := NewGatewayContainer(configContainer)
	useCaseContainer := NewUseCaseContainer(repositoryContainer, gatewayContainer)
	middlewareContainer := NewMiddlewareContainer(useCaseContainer, repositoryContainer)

	return &RootContainer{
		ConfigContainer:     configContainer,
		RepositoryContainer: repositoryContainer,
		GatewayContainer:    gatewayContainer,
		UseCaseContainer:    useCaseContainer,
		MiddlewareContainer: middlewareContainer,
	}
}
