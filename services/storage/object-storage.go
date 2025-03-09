package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/rs/zerolog/log"
)

type S3Client struct {
	client             *s3.Client
	expectedBUcketName string
	region             string
}

func NewS3Client(config ObjectStorageConfig) *S3Client {
	cfg := aws.Config{
		Region: config.Region,
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     config.AccessKey,
				SecretAccessKey: config.SecretKey,
			}, nil
		}),
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = &config.Url
	})
	return &S3Client{
		client:             client,
		expectedBUcketName: config.BucketName,
		region:             config.Region,
	}

}

func (s *S3Client) Ping() {
	_, err := s.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: &s.expectedBUcketName,
	})
	if err != nil {
		log.Warn().Caller().Msgf("bucket %s does not exist creating it", s.expectedBUcketName)
		_, err = s.client.CreateBucket(context.Background(), &s3.CreateBucketInput{
			Bucket: &s.expectedBUcketName,
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(s.region),
			},
		})
		if err != nil {
			log.Fatal().Err(err).Caller().Msg("unable to create bucket " + s.expectedBUcketName)
		}
	} else {
		log.Debug().Caller().Msgf("bucket %s exist", s.expectedBUcketName)
	}

}

func (s *S3Client) Close() error {
	return nil

}
