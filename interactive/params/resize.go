package params

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmbracelet/huh"
)

type Resize struct {
	fields     []huh.Field
	height     *string
	width      *string
	resizeType *string
	enlarge    *[]string
	extent     *[]string
}

func NewResize() *Resize {
	rs := &Resize{
		fields: []huh.Field{
			huh.NewInput().
				Key("resize.height").
				Description("The height of the image.").
				Title("Height"),
			huh.NewInput().
				Key("resize.width").
				Description("The width of the image.").
				Title("Width"),
			huh.NewSelect[string]().
				Title("Resize Type").
				Options(
					huh.NewOption("fit", "fit"),
					huh.NewOption("fill", "fill"),
					huh.NewOption("fill-down", "fill-down"),
					huh.NewOption("force", "force"),
					huh.NewOption("auto", "auto"),
				).WithKeyMap(huh.NewDefaultKeyMap()),
			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("true", "true"),
				).
				Description("Whether to enlarge the image.").
				WithKeyMap(
					huh.NewDefaultKeyMap(),
				),
			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("true", "true"),
				).
				Description("Whether to extend the image.").
				WithKeyMap(
					huh.NewDefaultKeyMap(),
				),
		},
	}

	rs.height = aws.String("")
	rs.width = aws.String("")
	rs.enlarge = &[]string{}
	rs.extent = &[]string{}
	rs.resizeType = aws.String("auto")

	rs.fields[0].(*huh.Input).Value(rs.height)
	rs.fields[1].(*huh.Input).Value(rs.width)
	rs.fields[2].(*huh.Select[string]).Value(rs.resizeType)
	rs.fields[3].(*huh.MultiSelect[string]).Value(rs.enlarge)
	rs.fields[4].(*huh.MultiSelect[string]).Value(rs.extent)

	return rs
}

func (r Resize) Display() string {
	return "resize"
}

func (r Resize) Value() string {
	if *r.height == "" || *r.width == "" {
		return ""
	}

	resize := "0"
	if *r.enlarge != nil && len(*r.enlarge) > 0 {
		resize = "1"
	}

	extent := "0"
	if *r.extent != nil && len(*r.extent) > 0 {
		extent = "1"
	}

	return *r.resizeType + ":" + *r.height + ":" + *r.width + ":" + resize + ":" + extent
}

func (r Resize) Key() string {
	return "rs"
}

func (r Resize) Input() []huh.Field {
	return r.fields
}
