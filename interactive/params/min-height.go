package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type MinHeight struct {
	paramValue *string
	field      *huh.Input
}

func NewMinHeight() *MinHeight {
	mh := &MinHeight{
		paramValue: nil,
		field: huh.NewInput().
			Key("min-height").
			Description("The minimum height of the image.").
			Title("Min Height"),
	}

	mh.paramValue = aws.String("")
	mh.field.Value(mh.paramValue)

	return mh
}

func (mh MinHeight) Display() string {
	return "min-height"
}

func (mh MinHeight) Value() string {
	return *mh.paramValue
}

func (mh MinHeight) Key() string {
	return "mh"
}

func (mh MinHeight) Input() []huh.Field {
	return []huh.Field{mh.field}
}
