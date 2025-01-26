package commands

import (
	"fmt"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
)

func GenerateCommand() *cli.Command {
	return &cli.Command{
		Name:   "generate",
		Usage:  "Generate an image from a URL",
		Action: runGenerate,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "image",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "Convert the image to the specified format",
			},
			&cli.StringSliceFlag{
				Name:    "params",
				Aliases: []string{"p"},
				Usage:   "Processing options to be passed to the generator ref: https://docs.imgproxy.net/usage/processing",
			},
		},
	}
}

func runGenerate(c *cli.Context) error {
	p := printer.NewPrinter()

	_, err := config.NewConfig(c.String("config"))
	if err != nil {
		return err
	}

	url, err := signURL(c.String("image"), c.StringSlice("params"), c.String("format"))
	if err != nil {
		return err
	}

	p.LogInfo("Url Generated...")

	println(url)

	return nil
}

func signURL(file string, params []string, formatS string) (string, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return "", fmt.Errorf("config not loaded")
	}

	ig, err := generator.NewGenerator(cfg)
	if err != nil {
		return "", err
	}

	format, err := ig.StringToFormat(formatS)
	if err != nil {
		return "", err
	}

	url, err := ig.GenerateUrl(file, params, format)

	if err != nil {
		return "", err
	}

	return url, nil
}
