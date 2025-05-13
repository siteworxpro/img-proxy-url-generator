package interactive

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	fieldStyle := lipgloss.NewStyle().PaddingBottom(1)

	if m.err != nil {
		return lipgloss.
			NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("red")).
			Background(lipgloss.Color("black")).
			Render("Error: " + m.err.Error() + "\n" + "Press Ctrl+C to exit.")
	}

	output := lipgloss.NewStyle().
		Foreground(lipgloss.Color("5")).
		PaddingBottom(1).
		Bold(true).
		Underline(true).
		Render("Welcome to the img-proxy URL Generator!") + "\n"

	for _, field := range m.Fields {
		output += fieldStyle.Render("\n" + field.View())
	}

	for _, field := range *m.selectedParams {
		for _, f := range field.Input() {
			output += fieldStyle.Render("\n" + f.View())
		}
	}

	if *m.url == "" {
		return output + "\n" + m.renderHelp()
	}

	// Generate the URL
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

	output += m.renderHelp()

	return output
}

func (m Model) renderHelp() string {
	return m.help.ShortHelpView([]key.Binding{
		key.NewBinding(key.WithKeys("ctrl+c", "esc"), key.WithHelp("esc", "quit")),
		key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "cursor up")),
		key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "cursor down")),
		key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "select")),
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
	})
}
