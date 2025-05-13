package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type MinWidth struct {
	paramValue *string
	field      *huh.Input
}

func NewMinWidth() *MinWidth {
	mw := &MinWidth{
		paramValue: nil,
		field: huh.NewInput().
			Key("min-width").
			Description("The minimum width of the image.").
			Title("Min Width"),
	}

	mw.paramValue = aws.String("")
	mw.field.Value(mw.paramValue)

	return mw
}

func (mw MinWidth) Display() string {
	return "min-width"
}

func (mw MinWidth) Value() string {
	return *mw.paramValue
}

func (mw MinWidth) Key() string {
	return "mw"
}

func (mw MinWidth) Input() []huh.Field {
	return []huh.Field{mw.field}
}
