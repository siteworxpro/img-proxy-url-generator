package generator

import "testing"

func TestGenerator_GenerateUrl(t *testing.T) {
	g, err := NewGenerator(Config{})
	if err != nil {
		t.Fatal(err)
	}

	url, _ := g.generatePlainUrl("local:///test.png")

	if url != "plain/local:///test.png" {
		t.Errorf("url is %s", url)
	}
}
