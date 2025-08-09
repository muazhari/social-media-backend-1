package container

import "social-media-backend-1/internal/outers/configs"

type ConfigContainer struct {
	OneDatastoreConfig  *configs.OneDatastoreConfig
	TwoDatastoreConfig  *configs.TwoDatastoreConfig
	FourDatastoreConfig *configs.FourDatastoreConfig
	FiveDatastoreConfig *configs.FiveDatastoreConfig
	AuthConfig          *configs.AuthConfig
}

func NewConfigContainer() *ConfigContainer {
	return &ConfigContainer{
		OneDatastoreConfig:  configs.NewOneDatastoreConfig(),
		TwoDatastoreConfig:  configs.NewTwoDatastoreConfig(),
		FourDatastoreConfig: configs.NewFourDatastoreConfig(),
		FiveDatastoreConfig: configs.NewFiveDatastoreConfig(),
		AuthConfig:          configs.NewAuthConfig(),
	}
}
