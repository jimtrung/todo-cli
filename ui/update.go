package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimtrung/todo-cli/model"
	"github.com/jimtrung/todo-cli/storage"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.State {

	case StateAdd, StateUpdate:
		var cmd tea.Cmd
		m.Input, cmd = m.Input.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				text := m.Input.Value()
				if text != "" {
					if m.State == StateAdd {
						m.Todos = append(m.Todos, model.Todo{Description: text})
					} else if m.State == StateUpdate && m.Cursor < len(m.Todos) {
						m.Todos[m.Cursor].Description = text
					}
				}
				m.State = StateList
			case tea.KeyEsc:
				m.State = StateList
			}
		}

		return m, cmd

	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit

			case "up":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down":
				if m.Cursor < len(m.Todos)-1 {
					m.Cursor++
				}
			case " ":
				m.Todos[m.Cursor].IsDone = !m.Todos[m.Cursor].IsDone
			case "1":
				m.State = StateAdd
				m.Input.SetValue("")
				return m, nil
			case "2":
				if len(m.Todos) > 0 {
					m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
					if m.Cursor >= len(m.Todos) && m.Cursor > 0 {
						m.Cursor--
					}
				}
			case "3":
				if len(m.Todos) > 0 {
					m.State = StateUpdate
					m.Input.SetValue(m.Todos[m.Cursor].Description)
				}
			}
		}
	}
	storage.SaveTodos(m.Todos)

	return m, nil
}
