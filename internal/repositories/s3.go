package repositories

import (
	"bytes"
	"context"
	"fireteam-core/internal/aws/awsclient"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type FileRepository interface {
	Put(fileName string, data []byte) error
	Get(key string) ([]byte, error)
}

type S3FileRepository struct {
	s3         *s3.Client
	bucketname string
	context    context.Context
}

func NewFileRepository(context context.Context, bucketname string) (FileRepository, error) {
	s3, err := awsclient.GetS3Client(context)
	if err != nil {
		return nil, fmt.Errorf("error when initialising connection to S3: %w", err)
	}

	return &S3FileRepository{
		s3:         s3,
		bucketname: bucketname,
		context:    context,
	}, nil
}

// Save data into a file S3
func (r *S3FileRepository) Put(fileName string, data []byte) error {
	params := &s3.PutObjectInput{
		Bucket: aws.String(r.bucketname),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(data),
	}
	if _, err := r.s3.PutObject(r.context, params); err != nil {
		return fmt.Errorf("error when writing object to S3: %w", err)
	}
	return nil
}

// Get data from a file in S3
func (r *S3FileRepository) Get(fileName string) ([]byte, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(r.bucketname),
		Key:    aws.String(fileName),
	}
	resp, err := r.s3.GetObject(r.context, params)
	if err != nil {
		return nil, fmt.Errorf("error when getting object from S3: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error when reading data from S3 object: %w", err)
	}
	return data, nil
}
