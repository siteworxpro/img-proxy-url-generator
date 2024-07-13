package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Config struct {
	AwsKey    string
	AwsSecret string
	AwsRole   string
	Bucket    string
}

type Service struct {
	s3     *s3.S3
	bucket string
}

func NewClient(config *Config) *Service {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.AwsKey, config.AwsSecret, ""),
		Region:      aws.String("us-east-1"),
	}))

	assumeRoleCredentials := stscreds.NewCredentials(awsSession, config.AwsRole)

	return &Service{
		s3:     s3.New(awsSession, &aws.Config{Credentials: assumeRoleCredentials}),
		bucket: config.Bucket,
	}
}
