package interactive

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	output := lipgloss.NewStyle().
		Foreground(lipgloss.Color("5")).
		Background(lipgloss.Color("0")).
		Bold(true).
		Underline(true).
		Render("Welcome to the img-proxy URL Generator!") + "\n\n"

	if m.err != nil {
		return output + "Error: " + m.err.Error() + "\n" + "Press Ctrl+C to exit.\n"
	}

	for _, field := range m.Fields {
		output += field.View() + "\n\n"
	}

	for _, field := range *m.selectedParams {
		for _, f := range field.Input() {
			output += f.View() + "\n\n"
		}
	}

	if *m.url == "" {
		return output + help()
	}

	params := make([]string, 0)
	for _, field := range *m.selectedParams {
		if field.Value() != "" {
			params = append(params, field.Key()+":"+field.Value())
		}
	}

	url, _ := m.generator.GenerateUrl(*m.url, params, *m.format)

	output += fmt.Sprintf("\nGenerated URL: %s\n\n", lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#0000ff")).
		Render(url))

	output += help()

	return output
}

func help() string {
	return lipgloss.
		NewStyle().
		Inline(true).
		Foreground(lipgloss.Color("#123456")).
		Render("Ctrl+C | Esc: exit * Tab: Next Field * Sft+Tab: Prev Field * Space: Select\n")
}
