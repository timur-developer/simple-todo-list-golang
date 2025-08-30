package tasks

type TaskData struct {
	tasks []Task
}

func CreateTaskData() TaskData {
	taskData := TaskData{}
	return taskData
}

func (td *TaskData) AddTask(task Task) {
	td.tasks = append(td.tasks, task)
}

func (td *TaskData) GetAllTasks() []Task {
	return td.tasks
}
