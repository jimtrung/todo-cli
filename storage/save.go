package storage

import (
	"encoding/json"
	"os"

	"github.com/jimtrung/todo-cli/model"
)

func SaveTodos(todos []model.Todo) error {
	/* Create new file if not exist */
	file, err := os.Create("todos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	/* Encode */
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(todos)
}
