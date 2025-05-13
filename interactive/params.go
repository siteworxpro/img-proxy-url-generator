package interactive

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

/**
 * Height
 */

type Height struct {
	paramValue *string
	field      *huh.Input
}

func NewHeight() *Height {
	h := &Height{
		paramValue: aws.String(""),
		field: huh.NewInput().
			Key("h").
			Description("The height of the image.").
			Title("Height"),
	}

	h.field.Value(h.paramValue)

	return h
}

func (h Height) display() string {
	return "height"
}

func (h Height) value() string {
	return *h.paramValue
}

func (h Height) key() string {
	return "h"
}

func (h Height) Input() huh.Field {
	return h.field
}

type Width struct {
	paramValue *string
	field      *huh.Input
}

/**
 * Width
 */

func NewWidth() *Width {
	w := &Width{
		paramValue: aws.String(""),
		field: huh.NewInput().
			Key("h").
			Description("The width of the image.").
			Title("With"),
	}

	w.field.Value(w.paramValue)

	return w
}

func (h Width) display() string {
	return "width"
}

func (h Width) value() string {
	return *h.paramValue
}

func (h Width) key() string {
	return "h"
}

func (h Width) Input() huh.Field {
	return h.field
}
