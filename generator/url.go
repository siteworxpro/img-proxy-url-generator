package generator

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"git.s.int/rrise/imgproxy/printer"
	"io"
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
			return nil, fmt.Errorf("Key expected to be hex-encoded string")
		}
	}

	return gen, nil
}

func (g *Generator) generatePlainUrl(file []byte) {

}

func (g *Generator) generateBase64Url(file []byte) string {
	return base64.RawURLEncoding.EncodeToString(file)
}

func pkcs7pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func (g *Generator) generateBaseAesEncUrl(file []byte) (string, error) {
	c, err := aes.NewCipher(g.config.encryptionKeyBin)
	if err != nil {
		return "", err
	}

	data := pkcs7pad(file, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	encryptedURL := g.generateBase64Url(ciphertext)

	return "enc/" + encryptedURL, nil
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
		url = "plain/" + file
	} else if g.config.encryptionKeyBin != nil {
		url, err = g.generateBaseAesEncUrl([]byte(file))
		if err != nil {
			return "", err
		}
	} else {
		url = g.generateBase64Url([]byte(file))
	}

	path := fmt.Sprintf("%s%s", paramString, url)

	if format != DEF {
		path = path + "." + string(format)
	}

	var signature string
	if len(g.config.keyBin) == 0 || len(g.config.saltBin) == 0 {
		signature = "insecure"

		printer.NewPrinter().LogWarning("Insecure url generated. Provide salt and key to sign and secure url.")

	} else {
		mac := hmac.New(sha256.New, g.config.keyBin)
		mac.Write(g.config.saltBin)
		mac.Write([]byte(path))
		signature = base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	}

	return fmt.Sprintf("%s/%s%s", g.config.Host, signature, path), nil
}
