package configs

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

type OneDatastoreConfig struct {
	Client *redis.Client
}

func NewOneDatastoreConfig() *OneDatastoreConfig {
	ctx := context.Background()
	option := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DATASTORE_1_HOST"), os.Getenv("DATASTORE_1_PORT")),
		Password: os.Getenv("DATASTORE_1_PASSWORD"),
		DB:       0,
	}
	client := redis.NewClient(option)

	pong, pingErr := client.Ping(ctx).Result()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println(pong)

	oneDatabaseConfig := &OneDatastoreConfig{
		Client: client,
	}

	return oneDatabaseConfig
}
