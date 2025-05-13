package params

import (
	"github.com/charmbracelet/huh"
)

type Enlarge struct {
	enlarge *[]string
	field   huh.Field
}

func NewEnlarge() *Enlarge {
	e := &Enlarge{
		enlarge: &[]string{},
		field: huh.NewMultiSelect[string]().
			Options(
				huh.NewOption("true", "true"),
			).
			Description("Whether to enlarge the image.").
			WithKeyMap(
				huh.NewDefaultKeyMap(),
			),
	}

	e.enlarge = &[]string{}
	e.field.(*huh.MultiSelect[string]).Value(e.enlarge)

	return e
}

func (e Enlarge) Display() string {
	return "enlarge"
}

func (e Enlarge) Value() string {
	value := "0"

	if len(*e.enlarge) > 0 {
		value = "1"
	}

	return value
}

func (e Enlarge) Key() string {
	return "el"
}

func (e Enlarge) Input() []huh.Field {
	return []huh.Field{e.field}
}
