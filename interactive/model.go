package interactive

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/urfave/cli/v2"
)

type Model struct {
	Form           *huh.Form
	generator      *generator.Generator
	url            *string
	format         *generator.Format
	selectedParams *[]UrlParam
	err            error
}

type UrlParam interface {
	value() string
	key() string
	Input() *huh.Input
}

func InitialModel(c *cli.Context) Model {

	m := Model{
		url:            aws.String(""),
		selectedParams: &[]UrlParam{},
		format:         generator.ToPtr(generator.DEF),
	}

	options := []UrlParam{
		Height{
			paramValue: "40",
		},
		Width{
			paramValue: "40",
		},
	}

	var huhOptions []huh.Option[UrlParam]
	for _, option := range options {
		huhOptions = append(huhOptions, huh.NewOption[UrlParam](option.value(), option))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("imgUrl").
				Description("The URL of the image to generate a proxy for.").
				Title("Image URL").
				Value(m.url).
				Prompt("Enter the image URL:"),

			huh.NewMultiSelect[UrlParam]().
				Description("Params to add to the URL.").
				Options(huhOptions...).
				Value(m.selectedParams),
			huh.NewSelect[generator.Format]().
				Description("Convert the image format.").
				Options(
					huh.NewOption[generator.Format]("JPEG", generator.JPG),
					huh.NewOption[generator.Format]("PNG", generator.PNG),
					huh.NewOption[generator.Format]("BMP", generator.BMP),
					huh.NewOption[generator.Format]("Default", generator.DEF),
				).
				Key("format").
				Title("Format").
				Value(m.format),
		),
		m.selectedOptionFields(),
	).WithShowHelp(false)

	m.Form = form

	cfg, _ := config.NewConfig(c.String("config"))

	if cfg == nil {
		err := fmt.Errorf("config not loaded")
		m.err = err

		return m
	}

	g, err := generator.NewGenerator(cfg)

	if err != nil {
		m.err = err
		return m
	}

	m.generator = g

	return m
}

func (m Model) selectedOptionFields() *huh.Group {
	fields := make([]huh.Field, 0)

	if m.selectedParams == nil || len(*m.selectedParams) == 0 {
		return huh.NewGroup(
			huh.NewText().Description("No params selected.").Title("Params").CharLimit(0),
		)
	}

	for _, param := range *m.selectedParams {
		fields = append(fields, param.Input())
	}

	return huh.NewGroup(fields...)
}

func (m Model) Init() tea.Cmd {
	return nil
}
