package config

import (
	"github.com/bigkevmcd/go-configparser"
	"sync"
)

type Config struct {
	initializeOnce sync.Once
	Generator      *GeneratorConfig
	Aws            *awsConfig
	Redis          *redisConfig
}

var c *Config

func GetConfig() *Config {
	if c == nil {
		return nil
	}

	return c
}

// NewConfig returns a new Config struct
func NewConfig(path string) (*Config, error) {

	if path == "" {
		path = "imgproxy.cfg"
	}

	p, err := configparser.NewConfigParserFromFile(path)
	if err != nil {
		return nil, err
	}

	c = &Config{}

	gc, err := getGeneratorConfig(p)
	if err != nil {
		return nil, err
	}
	c.Generator = gc

	if p.HasSection("aws") {
		c.Aws = getAwsConfig(p)
	}

	if p.HasSection("redis") {
		c.Redis = getRedisConfig(p)
	}

	return c, nil
}
