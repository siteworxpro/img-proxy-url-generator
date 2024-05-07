package generator

import (
	"testing"
)

func TestGenerator_GenerateBase64(t *testing.T) {
	g, err := NewGenerator(Config{})
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
