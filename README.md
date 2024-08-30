# TaskWarden

**TaskWarden** is a command-line task management tool written in Go. It allows you to create, list, mark as done, and delete tasks. TaskWarden uses a simple file-based storage system to persist tasks between sessions.

## Features

- **Add a Task**: Create new tasks with a simple command.
- **List Tasks**: Display all tasks, showing their completion status.
- **Mark a Task as Done**: Mark a task as completed by its ID.
- **Delete a Task**: Remove a task by its ID.
- **Persistent Storage**: Tasks are saved to a `tasks.json` file and persist between runs.

## Installation

To use TaskWarden, you need to have [Go](https://golang.org/) installed on your system. Once you have Go set up, you can clone the repository and run the application directly.

```bash
git clone https://github.com/yourusername/TaskWarden.git
cd TaskWarden
```

## Usage

TaskWarden is a CLI tool with the following commands:

### 1. Add a Task

To add a new task, use the `add` command followed by the task description:

```bash
go run cmd/taskwarden/main.go add "Buy groceries"
```

This command will add the task "Buy groceries" to your task list.

### 2. List Tasks

To list all tasks, use the `list` command:

```bash
go run cmd/taskwarden/main.go list
```

This will display all tasks along with their IDs and completion status.

### 3. Mark a Task as Done

To mark a task as done, use the `done` command followed by the task ID:

```bash
go run cmd/taskwarden/main.go done 1
```

This will mark the task with ID 1 as completed.

### 4. Delete a Task

To delete a task, use the `delete` command followed by the task ID:

```bash
go run cmd/taskwarden/main.go delete 1
```

This will delete the task with ID 1 from your list.

## File-Based Storage

Tasks are stored in a file called `tasks.json` located in the root of the project directory. This file is automatically created and updated by TaskWarden as you add, complete, or delete tasks.

## Running Tests

To run the test suite, use the `go test` command. This will execute all tests in the `internal/task` package.

```bash
go test -v ./internal/task
```

The tests cover all the core functionalities of TaskWarden, including adding tasks, listing tasks, marking them as done, and deleting them.

## Contribution

Contributions are welcome! If you have any ideas for new features or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
