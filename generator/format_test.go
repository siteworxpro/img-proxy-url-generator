package generator

import (
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"testing"
)

func TestGenerator_StringToFormat(t *testing.T) {
	g, err := NewGenerator(&config.Config{
		Generator: &config.GeneratorConfig{
			Key:  []byte("2c90317177aa7a3c44fa6804bf9bf466930f36ac9262bfdae972e836a9f83d239fd6bcee0c91a29ada58cc7329c787f35d2309f0984f2fd315e2c27bac8ac247"),
			Salt: []byte("2777def3372a385f4aa7e62b2b431927"),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	formats := []string{"jpg", "jpeg", "png", "gif", "bmp", "webp", "ico", ""}

	for _, format := range formats {
		format, err := g.StringToFormat(format)
		if err != nil {
			t.Error(err)
		}

		if !isFormat(format) {
			t.Error("format not match")
		}
	}
}

func TestGenerator_StringToFormatError(t *testing.T) {
	g, err := NewGenerator(&config.Config{
		Generator: &config.GeneratorConfig{
			Key:  []byte("f2d1f37016b7d12ab27b25377f39fd84e2d3368472ff096261ce7ac3e8490af429d43803836ad6a42a3bd9fb859a38137173619cb00bcb6fe3870e3feab2b764"),
			Salt: []byte("919ed8813f76abfd42968b10e05258db"),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = g.StringToFormat("notaformat")
	if err == nil {
		t.Error("format did not return error")
	}
}

func isFormat(t interface{}) bool {
	switch t.(type) {
	case Format:
		return true
	default:
		return false

	}
}
