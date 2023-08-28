package awsclient

import (
	"context"
	"fireteam-core/internal/aws/awsconfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Client(ctx context.Context) (*s3.Client, error) {
	client, err := createClient(ctx, func(cfg aws.Config) interface{} {
		return s3.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	return client.(*s3.Client), nil
}

func createClient(ctx context.Context, factory func(cfg aws.Config) interface{}) (interface{}, error) {
	awsConf, err := awsconfig.GetAwsConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := factory(awsConf)
	return client, nil
}
