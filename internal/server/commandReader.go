package server

import (
	"fmt"
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
		i, err := MarkDone()
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("Task marked Done successfully (ID: %d)\n", i)
		return output
	case "list":
		tasks, err := List("tasks.json")
		if err != nil {
			return err.Error()
		}
		output := fmt.Sprintf("\n%s\n", tasks)
		return output
	case "list-in-progress":
		tasks, err := ListInProgress()
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("\n%s\n", tasks)
	case "list-done":
		tasks, err := ListDone()
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("\n%s\n", tasks)
	case "list-todo":
		tasks, err := ListTodo()
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("\n%s\n", tasks)
	default:
		return "Invalid command"
	}
}
