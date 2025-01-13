package server

import (
	"os"
	"testing"
)

func TestCommandReader(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"cmd", "list"}, "list"},
		{[]string{"cmd", "list", "todo"}, "list-todo"},
		{[]string{"cmd", "add", "task description"}, "add"},
		{[]string{"cmd", "update", "1", "new description"}, "update"},
	}

	for _, test := range tests {
		os.Args = test.args
		result := CommandReader()
		if result != test.expected {
			t.Errorf("CommandReader() = %v; want %v", result, test.expected)
		}
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		fileName string
	}{
		{"test_tasks.json"},
	}

	for _, test := range tests {
		_, err := ReadFile(test.fileName)
		if err != nil {
			t.Errorf("ReadFile() = %v; want nil", err)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"cmd", "add", "task description"}, 1},
	}

	for _, test := range tests {
		os.Args = test.args
		result, err := Add()
		if err != nil {
			t.Errorf("Add() = %v; want nil", err)
		}
		if result != test.expected {
			t.Errorf("Add() = %v; want %v", result, test.expected)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"cmd", "update", "1", "new description"}, 1},
	}

	for _, test := range tests {
		os.Args = test.args
		result, err := Update()
		if err != nil {
			t.Errorf("Update() = %v; want nil", err)
		}
		if result != test.expected {
			t.Errorf("Update() = %v; want %v", result, test.expected)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"cmd", "delete", "1"}, 1},
	}

	for _, test := range tests {
		os.Args = test.args
		result, err := Delete()
		if err != nil {
			t.Errorf("Delete() = %v; want nil", err)
		}
		if result != test.expected {
			t.Errorf("Delete() = %v; want %v", result, test.expected)
		}
	}
}

func TestMarkInProgress(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"cmd", "mark-in-progress", "1"}, 1},
	}

	for _, test := range tests {
		os.Args = test.args
		result, err := MarkInProgress()
		if err != nil {
			t.Errorf("MarkInProgress() = %v; want nil", err)
		}
		if result != test.expected {
			t.Errorf("MarkInProgress() = %v; want %v", result, test.expected)
		}
	}
}

func TestMarkDone(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"cmd", "mark-done", "1"}, 1},
	}

	for _, test := range tests {
		os.Args = test.args
		result, err := MarkDone()
		if err != nil {
			t.Errorf("MarkDone() = %v; want nil", err)
		}
		if result != test.expected {
			t.Errorf("MarkDone() = %v; want %v", result, test.expected)
		}
	}
}
