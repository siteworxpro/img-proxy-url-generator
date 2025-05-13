package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type Quality struct {
	value *string
	field *huh.Input
}

func NewQuality() *Quality {
	z := &Quality{
		value: aws.String(""),
		field: huh.NewInput().
			Key("q").
			Description("Quality of the image 0-100").
			Title("Quality"),
	}

	z.field.Value(z.value)

	return z
}

func (z *Quality) Value() string {
	return *z.value
}

func (z *Quality) Display() string {
	return "quality"
}

func (z *Quality) Key() string {
	return "q"
}

func (z *Quality) Input() []huh.Field {
	return []huh.Field{z.field}
}
