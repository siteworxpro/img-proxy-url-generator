package interactive

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Helper: find index of current focusField in a given field slice.
//
//goland:noinspection GoMixedReceiverTypes
func (m *Model) findCurrentFieldIndex(fields []huh.Field) int {
	for i, field := range fields {
		if field == m.focusField {
			return i
		}
	}
	return -1
}

// Helper: blur current, focus next (wraps to 0)
//
//goland:noinspection GoMixedReceiverTypes
func (m *Model) focusNextField(fields []huh.Field) {
	index := m.findCurrentFieldIndex(fields)
	next := (index + 1) % len(fields)
	m.focusField.Blur()
	m.focusField = fields[next]
	m.focusField.Focus()
}

// Helper: focus a specific field by index
//
//goland:noinspection GoMixedReceiverTypes
func (m *Model) focusFieldByIndex(fields []huh.Field, index int) {
	if index >= 0 && index < len(fields) {
		m.focusField.Blur()
		m.focusField = fields[index]
		m.focusField.Focus()
	}
}

// Helper: get all selected param fields as a flat slice
//
//goland:noinspection GoMixedReceiverTypes
func (m *Model) selectedParamFields() []huh.Field {
	var fields []huh.Field
	for _, param := range *m.selectedParams {
		fields = append(fields, param.Input()...)
	}
	return fields
}

//goland:noinspection GoMixedReceiverTypes
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model := m // copy, since Bubble Tea prefers value receivers, but we'll operate on &model internally

	if model.focusField == nil && len(model.Fields) > 0 {
		model.Fields[0].Focus()
		model.focusField = model.Fields[0]
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "tab":
			mainFields := model.Fields
			paramFields := model.selectedParamFields()

			if model.inParamsFields && len(paramFields) > 0 {
				index := model.findCurrentFieldIndex(paramFields)
				if index == -1 {
					break
				}
				if index == len(paramFields)-1 {
					// Last param field: cycle to main fields
					model.inParamsFields = false
					model.focusFieldByIndex(mainFields, 0)
				} else {
					model.focusNextField(paramFields)
				}
			} else {
				index := model.findCurrentFieldIndex(mainFields)
				if index == -1 {
					break
				}
				if index == len(mainFields)-1 && len(paramFields) > 0 {
					// Last main field & params exist: go to params
					model.inParamsFields = true
					model.focusFieldByIndex(paramFields, 0)
				} else {
					model.focusNextField(mainFields)
				}
			}
		case "ctrl+c", "esc":
			return model, tea.Quit
		case "enter":
			return model, nil
		default:
			if model.focusField != nil {
				md, cmd := model.focusField.(huh.Field).Update(msg)
				if md != nil {
					model.focusField = md.(huh.Field)
				}
				return model, cmd
			}
		}
	}
	return model, nil
}
