package repositories

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"social-media-backend-1/internal/outers/configs"
	"time"
)

type FileRepository struct {
	FourDatastoreConfig *configs.FourDatastoreConfig
}

func NewFileRepository(fourDatastoreConfig *configs.FourDatastoreConfig) *FileRepository {
	return &FileRepository{
		FourDatastoreConfig: fourDatastoreConfig,
	}
}

func (r *FileRepository) Delete(ctx context.Context, bucketName string, objectName string) error {
	err := r.FourDatastoreConfig.Client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) Upload(ctx context.Context, bucketName string, objectName string, file io.Reader, fileSize int64, contentType string) error {
	isExists, err := r.FourDatastoreConfig.Client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !isExists {
		err = r.FourDatastoreConfig.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	_, err = r.FourDatastoreConfig.Client.PutObject(ctx, bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (r *FileRepository) GetURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	presignedURL, err := r.FourDatastoreConfig.Client.PresignedGetObject(ctx, bucketName, objectName, time.Hour, url.Values{})
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
