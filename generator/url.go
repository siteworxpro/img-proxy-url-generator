package generator

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Generator struct {
	config Config
}

type Config struct {
	Salt             []byte
	saltBin          []byte
	Key              []byte
	keyBin           []byte
	Host             string
	EncryptionKey    *string
	encryptionKeyBin []byte
	PlainUrl         bool
}

var PathPrefix string

func NewGenerator(config Config) (*Generator, error) {
	var err error

	gen := new(Generator)
	gen.config = config

	if gen.config.keyBin, err = hex.DecodeString(string(gen.config.Key)); err != nil {
		return nil, err
	}

	if gen.config.saltBin, err = hex.DecodeString(string(gen.config.Salt)); err != nil {
		return nil, err
	}

	if gen.config.EncryptionKey != nil && *gen.config.EncryptionKey != "" {
		if gen.config.encryptionKeyBin, err = hex.DecodeString(*gen.config.EncryptionKey); err != nil {
			return nil, fmt.Errorf("key expected to be hex-encoded string")
		}
	}

	return gen, nil
}

func (g *Generator) GenerateUrl(file string, params []string, format Format) (string, error) {

	if params == nil || len(params) == 0 || params[0] == "" {
		params = []string{"raw:1"}
	} else {
		params = append(params, "sm:1")
	}

	if PathPrefix != "" {
		file = PathPrefix + file
	}

	paramString := "/" + strings.Join(params, "/") + "/"

	var url string
	var err error
	if g.config.PlainUrl {
		url, _ = g.generatePlainUrl(file)
	} else if g.config.encryptionKeyBin != nil {
		url, err = g.generateBaseAesEncUrl([]byte(file))
	} else {
		url, _ = g.generateBase64Url([]byte(file))
	}

	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s%s", paramString, url)

	if format != DEF {
		path = path + "." + string(format)
	}

	signature := g.generateSignature(path)

	return fmt.Sprintf("%s/%s%s", g.config.Host, signature, path), nil
}
