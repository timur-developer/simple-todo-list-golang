package tasks

type TaskData struct {
	tasks map[string]Task
}

func CreateTaskData() TaskData {
	taskData := TaskData{tasks: make(map[string]Task)}
	return taskData
}

func (td *TaskData) AddTask(task Task) {
	td.tasks[task.GetTaskName()] = task
}

func (td *TaskData) GetAllTasks() map[string]Task {
	return td.tasks
}
