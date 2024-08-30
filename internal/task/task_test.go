package task

import (
	"os"
	"testing"
)

func TestNewTaskManager(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "tasks_test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tm, err := NewTaskManager(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tm.tasks) != 0 {
		t.Errorf("Expected no tasks, got %d", len(tm.tasks))
	}
}

func TestTaskManagerOperations(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "tasks_test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tm, err := NewTaskManager(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Table-driven tests for adding tasks
	addTests := []struct {
		title    string
		expected int
	}{
		{"Task 1", 1},
		{"Task 2", 2},
	}

	for _, tt := range addTests {
		task := tm.AddTask(tt.title)
		if task.Title != tt.title {
			t.Errorf("Expected task title to be %s, got %s", tt.title, task.Title)
		}
		if task.ID != tt.expected {
			t.Errorf("Expected task ID to be %d, got %d", tt.expected, task.ID)
		}
	}

	// Table-driven tests for listing tasks
	listTests := []struct {
		expectedCount  int
		expectedTitles []string
	}{
		{2, []string{"Task 1", "Task 2"}},
	}

	for _, tt := range listTests {
		tasks := tm.ListTasks()
		if len(tasks) != tt.expectedCount {
			t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(tasks))
		}
		for i, task := range tasks {
			if task.Title != tt.expectedTitles[i] {
				t.Errorf("Expected task title to be %s, got %s", tt.expectedTitles[i], task.Title)
			}
		}
	}

	// Table-driven tests for marking tasks as done
	markDoneTests := []struct {
		taskID   int
		expected bool
		errMsg   string
	}{
		{1, true, "Task 1 should be marked as done"},
		{2, true, "Task 2 should be marked as done"},
		{999, false, "Non-existent task should return an error"},
	}

	for _, tt := range markDoneTests {
		err := tm.MarkTaskDone(tt.taskID)
		if (err == nil) != tt.expected {
			t.Errorf("MarkTaskDone(%d): expected %v, got error %v", tt.taskID, tt.expected, err)
		}

		if tt.expected && err == nil {
			for _, task := range tm.tasks {
				if task.ID == tt.taskID && !task.Done {
					t.Errorf(tt.errMsg)
				}
			}
		}
	}

	// Table-driven tests for deleting tasks
	deleteTests := []struct {
		taskID        int
		expectedCount int
		expectErr     bool
		errMsg        string
	}{
		{1, 1, false, "Expected 1 task remaining after deleting Task 1"},
		{2, 0, false, "Expected 0 tasks remaining after deleting Task 2"},
		{999, 0, true, "Expected an error for non-existent task deletion"},
	}

	for _, tt := range deleteTests {
		err := tm.DeleteTask(tt.taskID)
		if (err != nil) != tt.expectErr {
			t.Errorf("DeleteTask(%d): expected error %v, got %v", tt.taskID, tt.expectErr, err)
		}

		tasks := tm.ListTasks()
		if len(tasks) != tt.expectedCount {
			t.Errorf(tt.errMsg)
		}
	}
}
