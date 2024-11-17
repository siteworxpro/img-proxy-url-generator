package generator

import (
	"fmt"
	"strings"
)

type Format string

const (
	JPG  Format = "jpg"
	PNG  Format = "png"
	BMP  Format = "bmp"
	WEBP Format = "webp"
	GIF  Format = "gif"
	ICO  Format = "ico"
	DEF  Format = "default"
)

func (g *Generator) StringToFormat(string string) (Format, error) {
	s := strings.ToLower(string)
	switch s {
	case "jpg":
		return JPG, nil
	case "jpeg":
		return JPG, nil
	case "png":
		return PNG, nil
	case "bmp":
		return BMP, nil
	case "webp":
		return WEBP, nil
	case "gif":
		return GIF, nil
	case "ico":
		return ICO, nil
	case "def":
	case "default":
	case "":
		return DEF, nil
	}

	return "", fmt.Errorf("%s is not a valid format", string)
}
