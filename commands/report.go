package commands

import (
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/report"
	"github.com/urfave/cli/v2"
)

func ReportCommand() *cli.Command {
	return &cli.Command{
		Name:  "report",
		Usage: "Generate usage report",
		Action: func(c *cli.Context) error {
			cf, err := config.NewConfig(c.String("config"))
			if err != nil {
				return err
			}

			return report.Handle(cf)
		},
	}
}
