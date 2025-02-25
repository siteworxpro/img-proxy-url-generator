package generator

import (
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"testing"
)

func TestGenerator_GenerateUrl(t *testing.T) {
	g, err := NewGenerator(&config.Config{
		Generator: &config.GeneratorConfig{
			Key:      []byte("f2d1f37016b7d12ab27b25377f39fd84e2d3368472ff096261ce7ac3e8490af429d43803836ad6a42a3bd9fb859a38137173619cb00bcb6fe3870e3feab2b764"),
			Salt:     []byte("919ed8813f76abfd42968b10e05258db"),
			PlainUrl: true,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	url, _ := g.generatePlainUrl("local:///test.png")

	if url != "plain/local:///test.png" {
		t.Errorf("url is %s", url)
	}
}
