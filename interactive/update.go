package interactive

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.focusField == nil {
		m.Fields[0].Focus()
		m.focusField = m.Fields[0]
	}

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "tab":
			if m.focusField != nil {

				index := -1

				selectedParamFields := make([]huh.Field, 0)
				for _, field := range *m.selectedParams {
					for _, f := range field.Input() {
						selectedParamFields = append(selectedParamFields, f)
					}
				}

				if m.inParamsFields {
					for i, field := range selectedParamFields {
						if field == m.focusField {
							index = i
							break
						}
					}
				} else {
					for i, field := range m.Fields {
						if field == m.focusField {
							index = i
							break
						}
					}
				}

				// if the field is not found, return
				if index == -1 {
					return m, nil
				}

				// if the field is the last one, and we have params selected go to the param fields
				if !m.inParamsFields && index == len(m.Fields)-1 && len(selectedParamFields) > 0 {
					m.focusField.Blur()
					m.inParamsFields = true
					paramsFields := selectedParamFields
					m.focusField = paramsFields[0]
					m.focusField.Focus()

					// if the field is the last one, and we have params selected go to the first non params field
				} else if m.inParamsFields && index == len(selectedParamFields)-1 {
					m.focusField.Blur()
					m.inParamsFields = false
					m.focusField = m.Fields[0]
					m.focusField.Focus()

					// if not in the params fields and the field is the last one, go to the first one
				} else if index == len(m.Fields)-1 && !m.inParamsFields {
					m.focusField.Blur()
					m.focusField = m.Fields[0]
					m.focusField.Focus()
				} else {
					// otherwise, go to the next field
					m.focusField.Blur()
					if m.inParamsFields {
						m.focusField = selectedParamFields[index+1]
					} else {
						m.focusField = m.Fields[index+1]
					}
					m.focusField.Focus()
				}
			}
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			return m, nil
		default:
			if m.focusField != nil {
				md, cmd := m.focusField.(huh.Field).Update(msg)

				if md != nil {
					m.focusField = md.(huh.Field)
				}

				return m, cmd
			}
		}
	}

	return m, nil
}
