# task-tracker-cli

Task Tracker CLI is a command-line interface application for managing tasks. It allows you to add, update, delete, and list tasks, as well as mark them as in-progress or done.

## Features

- Add a new task
- Update an existing task
- Delete a task
- List all tasks
- List tasks by status (todo, in-progress, done)
- Mark a task as in-progress
- Mark a task as done

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/natnael-alemayehu/task-tracker-cli.git
    ```
2. Navigate to the project directory:
    ```sh
    cd task-tracker-cli
    ```
3. Build the project:
    ```sh
    go build -o task-tracker-cli ./cmd/api
    ```

## Usage

Run the CLI with the following commands:

### Add a Task
```sh
./task-tracker-cli add "task description"
```

### Update a Task
```sh
./task-tracker-cli update <task_id> "new description"
```

### Delete a Task
```sh 
./task-tracker-cli delete <task_id>
```

### List All tasks
```sh
./task-tracker-cli list
```

### List Tasks by status
```sh
./task-tracker-cli list-todo
./task-tracker-cli list-in-progress
./task-tracker-cli list-done
```

### Mark a Task as In-Progress
```sh
./task-tracker-cli mark-in-progress <task_id>
```

### Mark a Task as Done
```sh
./task-tracker-cli mark-done <task_id>
```

### Testing
To run the tests, use the following command: 
```sh
go test ./internal/server
```