package commands

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/siteworxpro/img-proxy-url-generator/interactive"
	"github.com/urfave/cli/v2"
)

func Interactive() *cli.Command {
	return &cli.Command{
		Name:  "interactive",
		Usage: "Start an interactive session",
		Action: func(c *cli.Context) error {
			p := tea.NewProgram(interactive.InitialModel(c), tea.WithAltScreen())

			if _, err := p.Run(); err != nil {
				return fmt.Errorf("error running interactive session: %w", err)
			}

			return nil
		},
	}
}
