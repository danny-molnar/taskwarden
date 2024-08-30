package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/danny-molnar/taskwarden/internal/task"
)

func main() {
	taskManager, err := task.NewTaskManager("tasks.json")
	if err != nil {
		fmt.Println("Error initializing TaskManager:", err)
		return
	}

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
			status := "pending"
			if task.Done {
				status = "done"
			}
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Title, status)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskwarden done <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		if err := taskManager.MarkTaskDone(id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Marked task %d as done\n", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskwarden delete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		if err := taskManager.DeleteTask(id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Deleted task %d\n", id)
	default:
		fmt.Println("Unknown command")
	}
}
