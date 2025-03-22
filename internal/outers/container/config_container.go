package container

import "social-media-backend-1/internal/outers/configs"

type ConfigContainer struct {
	OneDatabaseConfig *configs.OneDatabaseConfig
	TwoDatabaseConfig *configs.TwoDatabaseConfig
}

func NewConfigContainer() *ConfigContainer {
	return &ConfigContainer{
		OneDatabaseConfig: configs.NewOneDatabaseConfig(),
		TwoDatabaseConfig: configs.NewTwoDatabaseConfig(),
	}
}
