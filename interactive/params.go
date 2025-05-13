package interactive

import "github.com/charmbracelet/huh"

type Height struct {
	UrlParam
	paramValue string
}

func (h Height) value() string {
	return "height"
}

func (h Height) key() string {
	return "h"
}

func (h Height) Input() *huh.Input {
	return huh.NewInput().
		Key(h.key()).
		Description("The height of the image.").
		Title("Height").
		Value(&h.paramValue)
}

type Width struct {
	UrlParam
	paramValue string
}

func (h Width) value() string {
	return "width"
}

func (h Width) key() string {
	return "w"
}

func (h Width) Input() *huh.Input {
	return huh.NewInput().
		Key(h.key()).
		Description("The width of the image.").
		Title("Width").
		Value(&h.paramValue)
}
