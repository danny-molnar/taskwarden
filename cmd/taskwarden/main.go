package main

import (
	"fmt"
	"os"

	"github.com/danny-molnar/taskwarden/internal/task"
)

func main() {
	taskManager := task.NewTaskManager()

	if len(os.Args) < 2 {
		fmt.Println("Usage: taskwarden <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskwarden add <task>")
			return
		}
		newTask := taskManager.AddTask(os.Args[2])
		fmt.Printf("Added task: %d: %s\n", newTask.ID, newTask.Title)
	case "list":
		tasks := taskManager.ListTasks()
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%d: %s [done: %v]\n", task.ID, task.Title, task.Done)
		}
	default:
		fmt.Println("Unknown command")
	}
}
