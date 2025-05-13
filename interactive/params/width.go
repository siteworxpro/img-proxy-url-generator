package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type Width struct {
	paramValue *string
	field      *huh.Input
}

func NewWidth() *Width {
	w := &Width{
		paramValue: aws.String(""),
		field: huh.NewInput().
			Key("h").
			Description("The width of the image.").
			Title("Width"),
	}

	w.field.Value(w.paramValue)

	return w
}

func (h Width) Display() string {
	return "width"
}

func (h Width) Value() string {
	return *h.paramValue
}

func (h Width) Key() string {
	return "h"
}

func (h Width) Input() []huh.Field {
	return []huh.Field{h.field}
}
