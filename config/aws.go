package config

import "github.com/bigkevmcd/go-configparser"

type awsConfig struct {
	AwsKey    string
	AwsSecret string
	AwsToken  string
	AwsRegion string
	AwsBucket string
	AwsRole   string
}

func getAwsConfig(p *configparser.ConfigParser) *awsConfig {
	ac := &awsConfig{}
	ac.AwsKey, _ = p.Get("aws", "key")
	ac.AwsSecret, _ = p.Get("aws", "secret")
	ac.AwsToken, _ = p.Get("aws", "token")
	ac.AwsRegion, _ = p.Get("aws", "region")
	ac.AwsBucket, _ = p.Get("aws", "bucket")
	ac.AwsRole, _ = p.Get("aws", "role")

	return ac
}
