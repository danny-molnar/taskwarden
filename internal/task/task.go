package task

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskManager struct {
	tasks  []Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  []Task{},
		nextID: 0,
	}
}

func (tm *TaskManager) AddTask(title string) Task {
	tm.nextID++
	task := Task{
		ID:    tm.nextID,
		Title: title,
		Done:  false,
	}
	tm.tasks = append(tm.tasks, task)
	return task
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.tasks
}
