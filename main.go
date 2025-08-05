package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimtrung/todo-cli/storage"
	"github.com/jimtrung/todo-cli/ui"
)

func main() {
	input := textinput.New()
	input.Placeholder = "Task description"
	input.Focus()

	todos, err := storage.LoadTodos()
	if err != nil {
		log.Fatalf("Failed to load the data: %v", err)
	}

	m := ui.Model{
		Todos: todos,
		Input: input,
		State: 0,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
