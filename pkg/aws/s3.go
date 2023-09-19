package aws

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io"
)

type S3Client struct {
	client     *minio.Client
	log        *zap.Logger
	bucketName string
}

func (s *S3Client) checkIfBucketExists(ctx context.Context, bucketName string) (bool, error) {
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (s *S3Client) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	log := s.log.Named("Save").With(
		zap.String("objectName", name),
		zap.Int64("size", size))

	uploadInfo, err := s.client.PutObject(ctx, s.bucketName, name, file, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	log.Info("UploadInfo", zap.String("uploadKey", uploadInfo.Key))
	return nil
}
