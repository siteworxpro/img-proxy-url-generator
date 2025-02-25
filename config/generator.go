package config

import (
	"fmt"
	"github.com/bigkevmcd/go-configparser"
)

type GeneratorConfig struct {
	Salt          []byte
	Key           []byte
	Host          string
	EncryptionKey string
	PlainUrl      bool
}

func getGeneratorConfig(p *configparser.ConfigParser) (*GeneratorConfig, error) {
	var config string
	var err error

	gc := &GeneratorConfig{}
	if !p.HasSection("img-proxy") {
		return nil, fmt.Errorf("config error - [img-proxy] config required")
	}

	config, _ = p.Get("img-proxy", "key")
	gc.Key = []byte(config)

	config, _ = p.Get("img-proxy", "salt")
	gc.Salt = []byte(config)

	if config, err = p.Get("img-proxy", "host"); err != nil {
		return nil, err
	}
	gc.Host = config

	config, _ = p.Get("img-proxy", "plain-url")
	gc.PlainUrl = config == "true" || config == "1"

	config, _ = p.Get("img-proxy", "encryption-key")
	gc.EncryptionKey = config

	return gc, nil
}
