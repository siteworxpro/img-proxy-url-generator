package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type Zoom struct {
	zoom  *string
	field huh.Field
}

func NewZoom() *Zoom {
	z := &Zoom{
		zoom: aws.String(""),
		field: huh.NewInput().
			Key("z").
			Description("Percentage to zoom the image (1.4 == 140%) .").
			Title("Zoom"),
	}

	z.field.(*huh.Input).Value(z.zoom)

	return z
}

func (z *Zoom) Value() string {
	return *z.zoom
}

func (z *Zoom) Display() string {
	return "zoom"
}

func (z *Zoom) Key() string {
	return "z"
}

func (z *Zoom) Input() []huh.Field {
	return []huh.Field{z.field}
}
