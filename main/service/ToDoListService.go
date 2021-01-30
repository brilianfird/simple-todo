package service

import (
	"time"
	"todo/main/model/entity"
	"todo/main/model/web"
	"todo/main/repository"
)

func CreateToDo(title string, description string) web.ToDoRestWeb {

	t := entity.ToDo{
		Timestamp: time.Now(),
		Title: title,
		Cleared: false,
		Description: description,
	}

	return web.ToDoRestWeb(repository.SaveToDoList(t))
}

func GetAllToDo() []web.ToDoRestWeb {
	all := repository.GetAll()

	var toReturn []web.ToDoRestWeb

	for _, val := range all {
		toReturn = append(toReturn, web.ToDoRestWeb(val))
	}

	return toReturn
}