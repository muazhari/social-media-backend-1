package configs

import (
	"fmt"
	"os"
)

type FiveDatastoreConfig struct {
	Broker  string
	GroupID string
}

func NewFiveDatastoreConfig() *FiveDatastoreConfig {
	host := os.Getenv("DATASTORE_5_HOST")
	port := os.Getenv("DATASTORE_5_PORT")
	broker := fmt.Sprintf("%s:%s", host, port)
	groupID := "social-media-backend-1"

	return &FiveDatastoreConfig{Broker: broker, GroupID: groupID}
}
