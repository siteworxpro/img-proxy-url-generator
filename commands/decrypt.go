package commands

import (
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
)

func DecryptCommand() *cli.Command {
	return &cli.Command{
		Name:  "decrypt",
		Usage: "decrypt an image url contents",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			pr := printer.NewPrinter()
			cfg, err := config.NewConfig(c.String("config"))
			if err != nil {
				return err
			}

			ig, err := generator.NewGenerator(cfg)
			if err != nil {
				return err
			}

			plain, err := ig.Decrypt(c.String("url"))
			if err != nil {
				return err
			}

			pr.LogSuccess(plain)

			return nil
		},
	}
}
