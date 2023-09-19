package aws

import (
	"app/internal/model"
	"app/pkg/logger"
	"context"
	"fmt"
	"github.com/doxanocap/pkg/lg"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const contentType = "multipart/form-data"

type Services struct {
	S3 *S3Client

	cfg *model.Config
}

func InitServices(cfg *model.Config) *Services {
	s := &Services{
		cfg: cfg,
	}
	s3client, err := s.InitS3()
	if err != nil {
		lg.Warnf("aws: s3: %s", err)
	}
	s.S3 = s3client
	return s
}

func (s *Services) InitS3() (*S3Client, error) {
	ctx := context.Background()

	minioClient, err := minio.New(s.cfg.AWS.S3.EndpointUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(s.cfg.AWS.S3.AccessKey, s.cfg.AWS.S3.SecretKey, ""),
		Secure: s.cfg.AWS.S3.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to S3: %v", err)
	}
	s3Client := &S3Client{
		client:     minioClient,
		log:        logger.Log.Named("[STORAGE][S3]"),
		bucketName: s.cfg.AWS.S3.BucketName,
	}

	exists, err := s3Client.checkIfBucketExists(ctx, s3Client.bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to find bucket: %v", err)
	}
	if !exists {
		return nil, fmt.Errorf("such bucket does not exists")
	}

	return s3Client, nil
}
