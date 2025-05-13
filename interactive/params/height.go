package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

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

func (h Height) Display() string {
	return "height"
}

func (h Height) Value() string {
	return *h.paramValue
}

func (h Height) Key() string {
	return "h"
}

func (h Height) Input() []huh.Field {
	return []huh.Field{h.field}
}
