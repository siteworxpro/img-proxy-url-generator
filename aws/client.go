package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/siteworxpro/img-proxy-url-generator/config"
)

type Service struct {
	s3     *s3.S3
	bucket string
}

func NewClient(config *config.Config) *Service {

	var accessCredentials *credentials.Credentials

	staticCredentials := credentials.NewStaticCredentials(config.Aws.AwsKey, config.Aws.AwsSecret, config.Aws.AwsToken)
	awsSession := session.Must(session.NewSession(&aws.Config{
		Credentials: staticCredentials,
		Region:      aws.String("us-east-1"),
	}))

	if config.Aws.AwsRole != "" {
		assumeRoleCredentials := stscreds.NewCredentials(awsSession, config.Aws.AwsRole)
		accessCredentials = assumeRoleCredentials
	} else {
		accessCredentials = staticCredentials
	}

	return &Service{
		s3:     s3.New(awsSession, &aws.Config{Credentials: accessCredentials}),
		bucket: config.Aws.AwsBucket,
	}
}
