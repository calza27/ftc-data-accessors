package helpers

import (
	"context"
	"fireteam-core/internal/repositories"
	"fireteam-core/internal/services"
)

func InitS3Service(context context.Context, bucket string) (services.S3Service, error) {
	fileRepo, err := repositories.NewFileRepository(context, bucket)
	if err != nil {
		return nil, err
	}
	s3Service, err := services.NewS3ServiceImpl(fileRepo)
	if err != nil {
		return nil, err
	}
	return s3Service, nil
}
