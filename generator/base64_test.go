package generator

import (
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"testing"
)

func TestGenerator_GenerateBase64(t *testing.T) {
	g, err := NewGenerator(&config.Config{
		Generator: &config.GeneratorConfig{
			Key:  []byte("f2d1f37016b7d12ab27b25377f39fd84e2d3368472ff096261ce7ac3e8490af429d43803836ad6a42a3bd9fb859a38137173619cb00bcb6fe3870e3feab2b764"),
			Salt: []byte("919ed8813f76abfd42968b10e05258db"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	url, err := g.generateBase64Url([]byte("test://file"))
	if err != nil {
		return
	}

	if url != "dGVzdDovL2ZpbGU" {
		t.Errorf("url is wrong: %s", url)
	}
}
