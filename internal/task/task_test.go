package task

import "testing"

func TestAddTask(t *testing.T) {
	tm := NewTaskManager()

	title := "Test Task"
	task := tm.AddTask(title)

	if task.Title != title {
		t.Errorf("Expected task title to be %s, got %s", title, task.Title)
	}

	if task.ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", task.ID)
	}

	if task.Done != false {
		t.Errorf("Expected task to be not done, got %v", task.Done)
	}
}

func TestListTasks(t *testing.T) {
	tm := NewTaskManager() // We'll implement this later

	tm.AddTask("Task 1")
	tm.AddTask("Task 2")

	tasks := tm.ListTasks() // This function should return all tasks

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "Task 1" || tasks[1].Title != "Task 2" {
		t.Errorf("Expected tasks to be 'Task 1' and 'Task 2', got '%s' and '%s'", tasks[0].Title, tasks[1].Title)
	}
}
