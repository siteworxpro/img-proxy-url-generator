package generator

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

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
