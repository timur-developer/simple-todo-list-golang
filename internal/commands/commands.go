package commands

import (
	"TestProject/internal/eventsdata"
	"TestProject/internal/tasks"
	"fmt"
	"github.com/k0kubun/pp"
	"strings"
	"time"
)

func Add(td *tasks.TaskData, ed *eventsdata.EventData, event *eventsdata.Event, data []string) error {
	task := tasks.CreateTask()
	task.SetTaskCreationTime(time.Now())
	if len(data) < 2 {
		fmt.Println("Вы не передали название или описание задачи. Используйте help, чтобы увидеть правильный вид команды.")
		return nil
	}
	taskName := data[0]
	taskDescription := strings.Join(data[1:], " ")
	task.SetTaskName(taskName)
	task.SetTaskText(taskDescription)
	td.AddTask(task)
	fmt.Printf("Задача %v была успешно добавлена!\n", taskName)
	return nil
}

func List(td *tasks.TaskData) error {
	taskList := td.GetAllTasks()
	if len(taskList) != 0 {
		fmt.Println("Полный список задач:")
		for i, _ := range taskList {
			if _, err := pp.Println(taskList[i]); err != nil {
				return err
			}
		}
	} else {
		fmt.Println("Пока вы не добавили ни одной задачи.")
	}

	return nil
}

func Del(td *tasks.TaskData, ed *eventsdata.EventData, event *eventsdata.Event, data []string) error {
	event.SetEventCreationTime(time.Now())
	lenData := len(data)
	if lenData > 1 {
		errMsg := "Название задачи не может быть больше одного слова. Используйте help, чтобы увидеть правильный вид команды."
		fmt.Println(errMsg)
		event.SetEventDescription(errMsg)
		ed.AddEvent(*event)
		return nil
	} else if lenData == 0 {
		errMsg := "Вы не ввели название задачи. Используйте help, чтобы увидеть правильный вид команды."
		fmt.Println(errMsg)
		event.SetEventDescription(errMsg)
		ed.AddEvent(*event)
		return nil
	}

	taskNameDelete := data[0]
	taskList := td.GetAllTasks()
	if _, ok := taskList[taskNameDelete]; !ok {
		fmt.Printf("Задачи с названием %v не было найдено. Воспользуйтесь командой 'list' для просмотра существующих задач\n", taskNameDelete)
	} else {
		delete(taskList, taskNameDelete)
		fmt.Printf("Задача с названием %v была успешно удалена!\n", taskNameDelete)
	}
	ed.AddEvent(*event)
	return nil
}

func Done(td *tasks.TaskData, ed *eventsdata.EventData, event *eventsdata.Event, data []string) error {
	event.SetEventCreationTime(time.Now())
	lenData := len(data)
	if lenData > 1 {
		errMsg := "Название задачи не может быть больше одного слова. Используйте help, чтобы увидеть правильный вид команды."
		fmt.Println(errMsg)
		event.SetEventDescription(errMsg)
		ed.AddEvent(*event)
		return nil
	} else if lenData == 0 {
		errMsg := "Вы не ввели название задачи. Используйте help, чтобы увидеть правильный вид команды."
		fmt.Println(errMsg)
		event.SetEventDescription(errMsg)
		ed.AddEvent(*event)
		return nil
	}

	taskNameUpdate := data[0]
	taskList := td.GetAllTasks()
	if _, ok := taskList[taskNameUpdate]; !ok {
		fmt.Printf("Задачи с названием %v не было найдено.\n", taskNameUpdate)
	} else {
		task := taskList[taskNameUpdate]
		task.MakeTaskDone()
		task.SetTaskExecutionTime(time.Now())
		taskList[taskNameUpdate] = task
		fmt.Printf("Задача %v была выполнена.\n", taskNameUpdate)
	}
	ed.AddEvent(*event)
	return nil
}

func Event(ed *eventsdata.EventData) error {
	allEvents := ed.GetAllEvents()
	fmt.Println("Полный список событий:")
	for i, _ := range allEvents {
		if _, err := pp.Println(allEvents[i]); err != nil {
			return err
		}
	}
	return nil
}

func Help() {
	fmt.Println("Список доступных команд:")
	fmt.Println("'add {заголовок задачи из одного слова}  {текст задачи из одного или нескольких слов}' - добавляет задачу")
	fmt.Println("'list' - позволяет получить полный список задач")
	fmt.Println("'del {заголовок существующей задачи}' - удаляет задачу по её названию")
	fmt.Println("'done {заголовок существующей задачи}' - помечает задачу по её названию как выполненную")
	fmt.Println("'event' - позволяет получить список всех событий")
	fmt.Println("'exit' - позволяет завершить выполнение программы")
}
