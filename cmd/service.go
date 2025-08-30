package cmd

import (
	"TestProject/internal/commands"
	"TestProject/internal/eventsdata"
	"TestProject/internal/tasks"
	"TestProject/internal/utilities"
	"fmt"
	"strings"
	"time"
)

func ExecuteProgramm() error {
	taskList := tasks.CreateTaskData()
	eventList := eventsdata.CreateEventData()
	for {
		commandText, err := utilities.ReadStrings("Введите команду:", "неправильная команда")
		event := eventsdata.CreateEvent()
		event.SetEventInput(commandText)
		event.SetEventCreationTime(time.Now())
		if err != nil {
			event.SetEventDescription(err.Error())
			eventList.AddEvent(event)
			return err
		}
		if commandText == "" {
			errMsg := "Команда не может быть пустой строкой. Используйте help для просмотра доступных команд."
			fmt.Println(errMsg)
			event.SetEventDescription(errMsg)
			eventList.AddEvent(event)
			continue
		}
		commandFields := strings.Fields(commandText)
		if len(commandFields) > 1 {
			errMsg := "Команда не может больше одного слова. Используйте help для просмотра доступных команд."
			event.SetEventDescription(errMsg)
			eventList.AddEvent(event)
			fmt.Println(errMsg)
			continue
		}

		command := commandFields[0]
		switch command {
		case "add":
			eventList.AddEvent(event)
			if err = commands.Add(&taskList, &eventList); err != nil {
				return err
			}
		case "list":
			eventList.AddEvent(event)
			if err = commands.List(&taskList); err != nil {
				return err
			}
		case "del":
			eventList.AddEvent(event)
			if err = commands.Del(&taskList, &eventList); err != nil {
				return err
			}
		case "done":
			eventList.AddEvent(event)
			if err = commands.Done(&taskList, &eventList); err != nil {
				return err
			}
		case "event":
			eventList.AddEvent(event)
			if err = commands.Events(&eventList); err != nil {
				eventList.AddEvent(event)
				return err
			}
		case "exit":
			eventList.AddEvent(event)
			fmt.Println("Программа завершает свою работу.")
			return nil
		case "help":
			commands.Help()
		default:
			errMsg := "Такой команды нет. Используйте help для просмотра доступных команд."
			event.SetEventDescription(errMsg)
			eventList.AddEvent(event)
			fmt.Println(errMsg)
		}
	}
}
