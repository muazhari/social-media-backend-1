package container

import (
	"social-media-backend-1/internal/outers/deliveries/gateways"
)

type GatewayContainer struct {
	AuthGateway *gateways.AuthGateway
	EDFGateway  *gateways.EDFGateway
}

func NewGatewayContainer(configContainer *ConfigContainer, repositoryContainer *RepositoryContainer) *GatewayContainer {
	return &GatewayContainer{
		AuthGateway: gateways.NewAuthGateway(configContainer.AuthConfig),
		EDFGateway:  gateways.NewEDFGateway(repositoryContainer.AccountRepository),
	}
}
