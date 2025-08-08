package configs

import (
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FourDatastoreConfig struct {
	Client *minio.Client
}

func NewFourDatastoreConfig() *FourDatastoreConfig {
	host := os.Getenv("DATASTORE_4_HOST")
	port := os.Getenv("DATASTORE_4_PORT")
	endpoint := fmt.Sprintf("%s:%s", host, port)
	accessKeyID := os.Getenv("DATASTORE_4_ROOT_USER")
	secretAccessKey := os.Getenv("DATASTORE_4_ROOT_PASSWORD")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	return &FourDatastoreConfig{
		Client: minioClient,
	}
}
