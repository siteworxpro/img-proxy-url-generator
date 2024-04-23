package generator

import "encoding/base64"

func (g *Generator) generateBase64Url(file []byte) (string, error) {
	return base64.RawURLEncoding.EncodeToString(file), nil
}
