package container

import (
	"social-media-backend-1/internal/outers/deliveries/gateways"
)

type GatewayContainer struct {
	AuthGateway *gateways.AuthGateway
}

func NewGatewayContainer(configContainer *ConfigContainer) *GatewayContainer {
	return &GatewayContainer{
		AuthGateway: gateways.NewAuthGateway(configContainer.AuthConfig),
	}
}
