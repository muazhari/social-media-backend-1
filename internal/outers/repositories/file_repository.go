package repositories

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"social-media-backend-1/internal/outers/configs"
)

type FileRepository struct {
	FourDatastoreConfig *configs.FourDatastoreConfig
}

func NewFileRepository(fourDatastoreConfig *configs.FourDatastoreConfig) *FileRepository {
	return &FileRepository{
		FourDatastoreConfig: fourDatastoreConfig,
	}
}

func (r *FileRepository) Upload(ctx context.Context, bucketName string, objectName string, file io.Reader, fileSize int64, contentType string) error {
	options := minio.PutObjectOptions{
		ContentType: contentType,
	}
	_, err := r.FourDatastoreConfig.Client.PutObject(ctx, bucketName, objectName, file, fileSize, options)
	return err
}

func (r *FileRepository) GetURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	reqParams := url.Values{}
	presignedURL, err := r.FourDatastoreConfig.Client.PresignedGetObject(ctx, bucketName, objectName, 0, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
