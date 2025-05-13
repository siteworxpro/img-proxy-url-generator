package interactive

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	cbhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/interactive/params"
	"github.com/urfave/cli/v2"
)

type Model struct {
	Fields         []huh.Field
	generator      *generator.Generator
	url            *string
	format         *generator.Format
	selectedParams *[]UrlParam
	err            error
	focusField     huh.Field
	inParamsFields bool
	help           cbhelp.Model
}

type UrlParam interface {
	Value() string
	Display() string
	Key() string
	Input() []huh.Field
}

func InitialModel(c *cli.Context) Model {

	m := Model{
		url:            aws.String(""),
		selectedParams: &[]UrlParam{},
		format:         generator.ToPtr(generator.DEF),
	}

	fields := make([]huh.Field, 0)
	fields = append(fields, m.initialFields()...)

	m.Fields = fields

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

	m.help = cbhelp.New()

	return m
}

func (m Model) initialFields() []huh.Field {

	fields := make([]huh.Field, 0)

	options := []UrlParam{
		params.NewHeight(),
		params.NewWidth(),
		params.NewResize(),
		params.NewMinWidth(),
		params.NewMinHeight(),
		params.NewQuality(),
		params.NewZoom(),
		params.NewEnlarge(),
	}

	var huhOptions []huh.Option[UrlParam]
	for _, option := range options {
		huhOptions = append(huhOptions, huh.NewOption[UrlParam](option.Display(), option))
	}

	fields = append(fields,
		huh.NewInput().
			Key("imgUrl").
			Description("The URL of the image to generate a proxy for.").
			Title("Image URL").
			Value(m.url).
			Prompt("Enter the image URL:"),
	)
	fields = append(fields,
		huh.NewMultiSelect[UrlParam]().
			Description("Params to add to the URL.").
			Options(huhOptions...).
			Value(m.selectedParams).
			WithKeyMap(&huh.KeyMap{
				MultiSelect: huh.MultiSelectKeyMap{
					Up:     key.NewBinding(key.WithKeys("up"), key.WithHelp("up", "up")),
					Down:   key.NewBinding(key.WithKeys("down"), key.WithHelp("down", "down")),
					Toggle: key.NewBinding(key.WithKeys(" "), key.WithHelp(" ", "toggle")),
				},
			}),
	)

	fields = append(fields, huh.NewSelect[generator.Format]().
		Description("Convert the image format.").
		Options(
			huh.NewOption[generator.Format]("JPEG", generator.JPG),
			huh.NewOption[generator.Format]("PNG", generator.PNG),
			huh.NewOption[generator.Format]("BMP", generator.BMP),
			huh.NewOption[generator.Format]("WEBP", generator.WEBP),
			huh.NewOption[generator.Format]("GIF", generator.GIF),
			huh.NewOption[generator.Format]("ICO", generator.ICO),
			huh.NewOption[generator.Format]("Default", generator.DEF),
		).
		Key("format").
		Title("Format").
		Value(m.format).
		WithKeyMap(&huh.KeyMap{
			Select: huh.SelectKeyMap{
				Up:   key.NewBinding(key.WithKeys("up"), key.WithHelp("up", "up")),
				Down: key.NewBinding(key.WithKeys("down"), key.WithHelp("down", "down")),
			},
		}),
	)

	return fields
}

func (m Model) Init() tea.Cmd {
	return nil
}
