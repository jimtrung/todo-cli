package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimtrung/todo-cli/model"
)

/* Create state */
type state int

const (
	StateList state = iota
	StateAdd
	StateUpdate
	StateRemove
)

type Model struct {
	Todos      []model.Todo
	Cursor     int
	State      state
	Input      textinput.Model
	IsInputing bool
}

func (m Model) Init() tea.Cmd {
	if m.State == StateAdd || m.State == StateUpdate {
		return textinput.Blink
	}
	return nil
}
