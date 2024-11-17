package generator

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func pkcs7pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func (g *Generator) Decrypt(s string) (string, error) {
	c, err := aes.NewCipher(g.config.encryptionKeyBin)
	if err != nil {
		return "", err
	}

	decoded, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	iv := decoded[:aes.BlockSize]
	cryptText := decoded[aes.BlockSize:]
	cbc := cipher.NewCBCDecrypter(c, iv)

	cbc.CryptBlocks(cryptText, cryptText)

	return string(cryptText), err
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

	encryptedURL, err := g.generateBase64Url(ciphertext)
	if err != nil {
		return "", err
	}

	return "enc/" + encryptedURL, nil
}
