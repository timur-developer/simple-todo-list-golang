package commands

import (
	"TestProject/internal/eventsdata"
	"TestProject/internal/tasks"
	"TestProject/internal/utilities"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"strings"
	"time"
)

func Add(td *tasks.TaskData, ed *eventsdata.EventData) error {
	task := tasks.CreateTask()

	for {
		taskNameInput, err := utilities.ReadStrings("Введите название задачи:", "Ошибка при вводе названия задачи")
		if err != nil {
			return err
		}
		event := eventsdata.CreateEvent()
		event.SetEventInput(taskNameInput)
		event.SetEventCreationTime(time.Now())
		taskNameFields := strings.Fields(taskNameInput)
		if lenHeader := len(taskNameFields); lenHeader == 1 {
			taskName := taskNameFields[0]
			task.SetTaskName(taskName)
			ed.AddEvent(event)
			break
		} else {
			errMsg := "Заголовок должен быть одним непустым словом"
			fmt.Println(errMsg)
			event.SetEventDescription(errMsg)
			ed.AddEvent(event)
		}
	}

	taskTextInput, err := utilities.ReadStrings("Введите текст задачи:", "Ошибка при вводе текста задачи")
	if err != nil {
		return err
	}
	task.SetTaskText(taskTextInput)
	task.SetTaskCreationTime(time.Now())
	td.AddTask(task)
	fmt.Printf("Задача %v была добавлена!\n", task.GetTaskName())
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

func Del(td *tasks.TaskData, ed *eventsdata.EventData) error {
	var taskNameDelete string
	for {
		taskNameInput, err := utilities.ReadStrings("Введите название задачи:", "Ошибка при вводе названия задачи")
		event := eventsdata.CreateEvent()
		event.SetEventInput(taskNameInput)
		event.SetEventCreationTime(time.Now())
		if err != nil {
			event.SetEventDescription(err.Error())
			ed.AddEvent(event)
			return err
		}
		taskNameFields := strings.Fields(taskNameInput)
		if len(taskNameFields) == 1 {
			taskNameDelete = taskNameFields[0]
			ed.AddEvent(event)
			break
		} else {
			errMsg := "длина заголовка не может быть больше одного слова"
			event.SetEventDescription(errMsg)
			ed.AddEvent(event)
			fmt.Println(errMsg)
		}
	}

	taskList := td.GetAllTasks()
	if _, ok := taskList[taskNameDelete]; !ok {
		fmt.Printf("Задачи с названием %v не было найдено.\n", taskNameDelete)
	} else {
		delete(taskList, taskNameDelete)
		fmt.Printf("Задача с названием %v была успешно удалена!\n", taskNameDelete)
	}
	return nil
}

func Done(td *tasks.TaskData, ed *eventsdata.EventData) error {
	taskNameInput, err := utilities.ReadStrings("Введите название задачи:", "Ошибка при вводе названия задачи")
	event := eventsdata.CreateEvent()
	event.SetEventInput(taskNameInput)
	if err != nil {
		event.SetEventDescription(err.Error())
		ed.AddEvent(event)
		return err
	}
	taskNameFields := strings.Fields(taskNameInput)
	if len(taskNameFields) > 1 {
		errMsg := "длина заголовка не может быть больше одного слова"
		event.SetEventDescription(errMsg)
		ed.AddEvent(event)
		return errors.New(errMsg)
	}
	taskNameUpdate := taskNameFields[0]
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

	return nil
}

func Events(ed *eventsdata.EventData) error {
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
	fmt.Println("'add' -> {заголовок задачи из одного слова} -> {текст задачи из одного или нескольких слов} - добавляет задачу")
	fmt.Println("'list' - позволяет получить полный список задач")
	fmt.Println("'del' -> {заголовок существующей задачи} - удаляет задачу по её названию")
	fmt.Println("'done' -> {заголовок существующей задачи} - помечает задачу по её названию как выполненную")
	fmt.Println("'event' - позволяет получить список всех событий")
	fmt.Println("'exit' - позволяет завершить выполнение программы")
}
