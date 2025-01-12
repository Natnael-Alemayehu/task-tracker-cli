package server

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/natnael-alemayehu/task-tracker-cli/internal/data"
)

func ReadCommand() string {
	switch CommandReader() {
	case "add":
		id, err := Add()
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("Task added successfully (ID: %d)\n", id)
		return output
	case "update":
		id, err := Update()
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("Task updated successfully (ID: %d)\n", id)
		return output
	case "delete":
		id, err := Delete()
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("Task deleted successfully (ID: %d)\n", id)
		return output
	case "mark-in-progress":
		i, err := MarkInProgress()
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("Task marked in progress successfully (ID: %d)\n", i)
		return output
	case "mark-done":
		return MarkDone()
	case "list":
		return List()
	case "list-in-progress":
		return ListInProgress()
	case "list-done":
		return ListDone()
	default:
		return "Invalid command"
	}
}

func ReadFile(fileName string) (*os.File, error) {
	data, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return data, nil
}

func ListTaks() error {
	file, err := os.Open("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return err
	}
	return nil
}
