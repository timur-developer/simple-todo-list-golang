package tasks

import (
	"time"
)

type Task struct {
	name          string
	text          string
	CreationTime  time.Time
	isDone        bool
	ExecutionTime time.Duration
}

func CreateTask() Task {
	task := Task{}
	return task
}

func (t *Task) SetTaskName(name string) {
	t.name = name
}

func (t *Task) SetTaskText(text string) {
	t.text = text
}

func (t *Task) SetTaskCreationTime(time time.Time) {
	t.CreationTime = time
}

func (t *Task) SetTaskExecutionTime(time time.Duration) {
	t.ExecutionTime = time
}

func (t *Task) MakeTaskDone() {
	t.isDone = true
}

func (t *Task) GetTaskName() string {
	return t.name
}
func (t *Task) GetTaskText() string {
	return t.text
}

func (t *Task) GetTaskCreationTime() time.Time {
	return t.CreationTime
}

func (t *Task) GetTaskExecutionTime() time.Duration {
	return t.ExecutionTime
}

func (t *Task) GetTaskStatus() bool {
	return t.isDone
}
