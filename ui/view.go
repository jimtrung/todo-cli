package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var builder strings.Builder

	builder.WriteString("./todo-cli\n")

	/* Render todos */
	for i, todo := range m.Todos {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		check := " "
		if todo.IsDone {
			check = "x"
		}

		builder.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, check, todo.Description))
	}

	/* Other states */
	if m.State == StateAdd {
		builder.WriteString("\nAdding task:\n")
		builder.WriteString(m.Input.View())
	} else if m.State == StateUpdate {
		builder.WriteString("\nUpdating task:\n")
		builder.WriteString(m.Input.View())
	} else if m.State == StateRemove {
		/* Do somthing here ? */
	} else if m.State == StateList {
		builder.WriteString("\nActions:\n")
		builder.WriteString("[1] Add\n[2] Remove\n[3] Update\n")
	}

	return builder.String()
}
