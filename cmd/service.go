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
	taskData := tasks.CreateTaskData()
	eventData := eventsdata.CreateEventData()
	for {
		commandText, err := utilities.ReadStrings("Введите команду:", "неправильная команда")
		event := eventsdata.CreateEvent()
		event.SetEventInput(commandText)
		event.SetEventCreationTime(time.Now())
		if err != nil {
			event.SetEventDescription(err.Error())
			eventData.AddEvent(event)
			return err
		}
		if commandText == "" {
			errMsg := "Команда не может быть пустой строкой. Используйте help для просмотра доступных команд."
			fmt.Println(errMsg)
			event.SetEventDescription(errMsg)
			eventData.AddEvent(event)
			continue
		}
		commandFields := strings.Fields(commandText)

		command := commandFields[0]
		textAfterCommand := commandFields[1:]
		switch command {
		case "add":
			if err = commands.Add(&taskData, &eventData, &event, textAfterCommand); err != nil {
				return err
			}
		case "list":
			eventData.AddEvent(event)
			if err = commands.List(&taskData); err != nil {
				return err
			}
		case "del":
			if err = commands.Del(&taskData, &eventData, &event, textAfterCommand); err != nil {
				return err
			}
		case "done":
			if err = commands.Done(&taskData, &eventData, &event, textAfterCommand); err != nil {
				return err
			}
		case "event":
			eventData.AddEvent(event)
			if err = commands.Event(&eventData); err != nil {
				eventData.AddEvent(event)
				return err
			}
		case "exit":
			eventData.AddEvent(event)
			fmt.Println("Программа завершает свою работу.")
			return nil
		case "help":
			commands.Help()
		default:
			errMsg := "Такой команды нет. Используйте help для просмотра доступных команд."
			event.SetEventDescription(errMsg)
			eventData.AddEvent(event)
			fmt.Println(errMsg)
		}
	}
}
