package services

import (
	"fireteam-core/internal/repositories"
	"fmt"
)

type S3Service interface {
	Put(fileName string, data []byte) error
	Get(fileName string) ([]byte, error)
}

type S3ServiceImpl struct {
	s3FileRepository *repositories.FileRepository
}

func NewS3ServiceImpl(fileRepo repositories.FileRepository) (S3Service, error) {
	if fileRepo == nil {
		return nil, fmt.Errorf("file repository is required to initilize s3 service")
	}
	return &S3ServiceImpl{
		s3FileRepository: &fileRepo,
	}, nil
}

// Put saves the data to the s3 bucket with the given file name
func (s *S3ServiceImpl) Put(fileName string, data []byte) error {
	err := (*s.s3FileRepository).Put(fileName, data)
	if err != nil {
		return fmt.Errorf("error saving data file - %s: %w", fileName, err)
	}
	return nil
}

// Get gets the data from the s3 bucket with the given file name
func (s *S3ServiceImpl) Get(fileName string) ([]byte, error) {
	data, err := (*s.s3FileRepository).Get(fileName)
	if err != nil {
		return nil, fmt.Errorf("error getting data for file - %s: %w", fileName, err)
	}
	return data, nil
}
