package server

import (
	"fmt"
	"os"
)

func CommandReader() string {
	word := os.Args[1]
	return word
}

func Add() string {
	word := os.Args[2]
	out := fmt.Sprintf("Task added successfully (ID: %s)", word)
	return out
}

func Update() string {
	id := os.Args[2]
	task := os.Args[3]
	out := fmt.Sprintf("id: %s, task: %s\n", id, task)
	return out
}

func Delete() string {
	id := os.Args[2]
	out := fmt.Sprintln(id)
	return out
}

func MarkInProgress() string {
	id := os.Args[2]
	out := fmt.Sprintln(id)
	return out
}

func MarkDone() string {
	id := os.Args[2]
	return fmt.Sprintln(id)
}
