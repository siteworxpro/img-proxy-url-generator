package interactive

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "tab":
			if m.Form.GetFocusedField().GetKey() == "format" {
				return m, m.Form.NextGroup()
			}

			c := m.Form.NextField()
			return m, c
		case "shift+tab":
			if m.Form.GetFocusedField().GetKey() == "imgUrl" {
				return m, nil
			}

			c := m.Form.PrevField()
			return m, c
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			return m, nil
		default:
			form, cmd := m.Form.Update(msg)
			m.Form = form.(*huh.Form)

			return m, cmd
		}
	}

	return m, nil
}
