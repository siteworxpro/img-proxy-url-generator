package generator

import "testing"

func TestGenerator_StringToFormat(t *testing.T) {
	g, err := NewGenerator(Config{})
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
	g, err := NewGenerator(Config{})
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
