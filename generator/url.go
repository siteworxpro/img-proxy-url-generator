package generator

import (
	"encoding/hex"
	"fmt"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"strings"
)

type Generator struct {
	keyBin        []byte
	salt          []byte
	encryptionKey []byte
}

var PathPrefix string

func NewGenerator(config *config.Config) (*Generator, error) {
	var err error

	gen := new(Generator)

	if gen.keyBin, err = hex.DecodeString(string(config.Generator.Key)); err != nil {
		return nil, err
	}

	if gen.salt, err = hex.DecodeString(string(config.Generator.Salt)); err != nil {
		return nil, err
	}

	if config.Generator.EncryptionKey != "" {
		if gen.encryptionKey, err = hex.DecodeString(config.Generator.EncryptionKey); err != nil {
			return nil, fmt.Errorf("key expected to be hex-encoded string")
		}
	}

	return gen, nil
}

func (g *Generator) GenerateUrl(file string, params []string, format Format) (string, error) {

	if params == nil || len(params) == 0 || params[0] == "" {
		params = []string{"raw:1"}
	}
	params = append(params, "sm:1")

	if PathPrefix != "" {
		file = PathPrefix + file
	}

	paramString := "/" + strings.Join(params, "/") + "/"

	var url string
	var err error
	if config.GetConfig().Generator.PlainUrl {
		url, _ = g.generatePlainUrl(file)
	} else if g.encryptionKey != nil {
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

	return fmt.Sprintf("%s/%s%s", config.GetConfig().Generator.Host, signature, path), nil
}
