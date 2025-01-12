package server

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/natnael-alemayehu/task-tracker-cli/internal/data"
)

func CommandReader() string {
	word := os.Args[1]
	return word
}

func Add() (int, error) {
	description := os.Args[2]

	file, err := ReadFile("tasks.json")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var tasks data.Tasks

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	if stat.Size() == 0 {
		tasks = data.Tasks{
			Tasks: []data.Task{},
		}
	} else {
		err = json.NewDecoder(file).Decode(&tasks)
		if err != nil {
			return 0, fmt.Errorf("error decoding tasks: %w", err)
		}
	}

	newTask := data.Task{
		ID:          len(tasks.Tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks.Tasks = append(tasks.Tasks, newTask)

	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return 0, err
	}

	return newTask.ID, nil
}

func Update() (int, error) {
	id_str := os.Args[2]
	description := os.Args[3]

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return 0, fmt.Errorf("error converting id to int: %w", err)
	}

	file, err := ReadFile("tasks.json")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}
	if stat.Size() == 0 {
		return 0, fmt.Errorf("no tasks to update")
	}

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return 0, fmt.Errorf("error decoding tasks: %w", err)
	}

	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks[i].Description = description
			tasks.Tasks[i].UpdatedAt = time.Now()
			break
		}

	}

	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Delete() (int, error) {
	id_str := os.Args[2]
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return 0, err
	}

	file, err := ReadFile("tasks.json")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	if stat.Size() == 0 {
		return 0, fmt.Errorf("no tasks to delete")
	}

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return 0, fmt.Errorf("error decoding tasks: %w", err)
	}
	fmt.Printf("Tasks: %v", tasks.Tasks)
	for i, task := range tasks.Tasks {
		if id == task.ID {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			fmt.Printf("New task: %v", tasks.Tasks)
			err = file.Truncate(0)
			if err != nil {
				return 0, err
			}
			file.Seek(0, 0)
			err = json.NewEncoder(file).Encode(&tasks)
			if err != nil {
				return 0, err
			}
			return id, nil
		}
	}
	return 0, fmt.Errorf("task with id %d not found", id)
}

func MarkInProgress() string {
	id := os.Args[2]
	out := fmt.Sprint(id)
	return out
}

func MarkDone() string {
	id := os.Args[2]
	return fmt.Sprint(id)
}

func List() string {
	return "List"
}

func ListInProgress() string {
	return "List in progress"
}

func ListDone() string {
	return "List done"
}

func ListTodo() string {
	return "List todo"
}
