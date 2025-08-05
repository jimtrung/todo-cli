package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jimtrung/todo-cli/model"
)

func LoadTodos() ([]model.Todo, error) {
	/* Open the file */
	file, err := os.Open("todos.json")
	if err != nil {
		// If empty file then return empty slice
		if errors.Is(err, os.ErrNotExist) {
			return []model.Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	// Check file size
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if info.Size() == 0 {
		// file exists but is empty
		return []model.Todo{}, nil
	}

	/* Decode to struct */
	var todos []model.Todo

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}
