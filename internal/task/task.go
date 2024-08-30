package task

import (
	"encoding/json"
	"errors"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskManager struct {
	tasks  []Task
	nextID int
	file   string
}

func NewTaskManager(file string) (*TaskManager, error) {
	tm := &TaskManager{
		tasks:  []Task{},
		nextID: 0,
		file:   file,
	}
	if err := tm.LoadTasks(); err != nil {
		return nil, err
	}
	return tm, nil
}

func (tm *TaskManager) AddTask(title string) Task {
	tm.nextID++
	task := Task{ID: tm.nextID, Title: title, Done: false}
	tm.tasks = append(tm.tasks, task)
	tm.SaveTasks()
	return task
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.tasks
}

func (tm *TaskManager) MarkTaskDone(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Done = true
			return tm.SaveTasks()
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) DeleteTask(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return tm.SaveTasks()
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) SaveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tm.file, data, 0644)
}

func (tm *TaskManager) LoadTasks() error {
	if _, err := os.Stat(tm.file); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(tm.file)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, &tm.tasks)
	if err != nil {
		return err
	}

	for _, task := range tm.tasks {
		if task.ID > tm.nextID {
			tm.nextID = task.ID
		}
	}
	return nil
}
