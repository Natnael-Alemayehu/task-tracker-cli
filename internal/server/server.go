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
	if word == "list" && len(os.Args) > 2 {
		word = os.Args[1] + "-" + os.Args[2]
		fmt.Println(word)
		return word
	}
	return word
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

	err = file.Truncate(0)
	if err != nil {
		return 0, err
	}
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

	err = file.Truncate(0)
	if err != nil {
		return 0, err
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

func MarkInProgress() (int, error) {
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
		return 0, fmt.Errorf("no tasks to mark in progress")
	}

	var tasks data.Tasks
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return 0, fmt.Errorf("error decoding tasks: %w", err)
	}

	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks[i].Status = "in-progress"
			tasks.Tasks[i].UpdatedAt = time.Now()
			break
		}
	}
	err = file.Truncate(0)
	if err != nil {
		return 0, err
	}
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func MarkDone() (int, error) {
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
		return 0, fmt.Errorf("no tasks to mark in progress")
	}

	var tasks data.Tasks
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return 0, fmt.Errorf("error decoding tasks: %w", err)
	}

	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks[i].Status = "done"
			tasks.Tasks[i].UpdatedAt = time.Now()
			break
		}
	}
	err = file.Truncate(0)
	if err != nil {
		return 0, err
	}
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func List(name string) ([]byte, error) {
	file, err := ReadFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err.Error() == "EOF" {
		return nil, fmt.Errorf("no tasks to list")
	} else if err != nil {
		return nil, err
	}

	t, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, err
	}

	return t, nil
}

func ListInProgress() ([]byte, error) {
	file, err := ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	var inProgressTasks data.Tasks
	for _, task := range tasks.Tasks {
		if task.Status == "in-progress" {
			inProgressTasks.Tasks = append(inProgressTasks.Tasks, task)
		}
	}

	t, err := json.MarshalIndent(inProgressTasks, "", "  ")
	if err != nil {
		return nil, err
	}

	return t, nil
}

func ListDone() ([]byte, error) {
	file, err := ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	var inProgressTasks data.Tasks
	for _, task := range tasks.Tasks {
		if task.Status == "done" {
			inProgressTasks.Tasks = append(inProgressTasks.Tasks, task)
		}
	}

	t, err := json.MarshalIndent(inProgressTasks, "", "  ")
	if err != nil {
		return nil, err
	}

	return t, nil
}

func ListTodo() ([]byte, error) {
	file, err := ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks data.Tasks

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	var inProgressTasks data.Tasks
	for _, task := range tasks.Tasks {
		if task.Status == "todo" {
			inProgressTasks.Tasks = append(inProgressTasks.Tasks, task)
		}
	}

	t, err := json.MarshalIndent(inProgressTasks, "", "  ")
	if err != nil {
		return nil, err
	}

	return t, nil
}
