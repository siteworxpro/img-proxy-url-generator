package main

import (
	"fmt"
	"github.com/bigkevmcd/go-configparser"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
	"os"
)

var keyBin, saltBin []byte

var imgGenerator *generator.Generator

var Version = "v0.0.0"

func main() {

	pr := printer.NewPrinter()

	app := &cli.App{
		Name:           "img-proxy-url-generator",
		Usage:          "URL Generator for the img proxy service",
		DefaultCommand: "generate",
		Version:        Version,
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generate an image from a URL",
				Action: func(c *cli.Context) error {
					return run(c, pr)
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "image",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Config file to load from",
				DefaultText: "imgproxy.cfg",
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
		Action: func(c *cli.Context) error {
			return run(c, pr)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		pr.LogError(err.Error())

		os.Exit(1)
	}
}

func run(c *cli.Context, p *printer.Printer) error {
	err := initGenerator(c.String("config"))

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

func initGenerator(config string) error {
	var err error

	if config == "" {
		config = "imgproxy.cfg"
	}

	p, err := configparser.NewConfigParserFromFile(config)
	if err != nil {
		return err
	}

	if !p.HasSection("img-proxy") {
		return fmt.Errorf("config error - [img-proxy] config required")
	}

	config, err = p.Get("img-proxy", "key")
	if config != "" {
		keyBin = []byte(config)
	}

	config, err = p.Get("img-proxy", "salt")
	saltBin = []byte(config)

	hostConf, err := p.Get("img-proxy", "host")
	if err != nil {
		return err
	}

	plainConfig, err := p.Get("img-proxy", "plain-url")

	encKey, err := p.Get("img-proxy", "encryption-key")

	generatorConfig := generator.Config{
		Salt:          saltBin,
		Key:           keyBin,
		Host:          hostConf,
		EncryptionKey: &encKey,
		PlainUrl:      plainConfig != "",
	}

	imgGenerator, err = generator.NewGenerator(generatorConfig)
	if err != nil {
		return err
	}

	return nil
}

func signURL(file string, params []string, formatS string) (string, error) {
	format, err := imgGenerator.StringToFormat(formatS)
	if err != nil {
		return "", err
	}

	url, err := imgGenerator.GenerateUrl(file, params, format)

	if err != nil {
		return "", err
	}

	return url, nil
}
