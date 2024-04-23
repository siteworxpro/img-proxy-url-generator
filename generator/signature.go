package generator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
)

func (g *Generator) generateSignature(path string) string {
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
	return signature
}
