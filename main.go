package main

import (
	cliCommands "github.com/siteworxpro/img-proxy-url-generator/commands"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
	"os"
)

var Version = "v0.0.0"

func main() {

	pr := printer.NewPrinter()

	var commands []*cli.Command
	commands = append(commands, cliCommands.GenerateCommand())
	commands = append(commands, cliCommands.ServerCommand())
	commands = append(commands, cliCommands.ReportCommand())
	commands = append(commands, cliCommands.GrpcCommand())
	commands = append(commands, cliCommands.DecryptCommand())
	commands = append(commands, cliCommands.Interactive())

	app := &cli.App{
		Name:           "img-proxy-url-generator",
		Usage:          "URL Generator for the img proxy service",
		DefaultCommand: "interactive",
		Version:        Version,
		Commands:       commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Config file to load from",
				DefaultText: "imgproxy.cfg",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		pr.LogError(err.Error())

		os.Exit(1)
	}
}
