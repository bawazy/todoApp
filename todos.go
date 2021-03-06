package main

import (
	"encoding/json"
	"io/ioutil"
)

type todo struct {
	Id        int
	Task      string
	Completed string
}

func getTodos() (todos []todo) {

	filebytes, err := ioutil.ReadFile("todos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(filebytes, &todos)
	if err != nil {
		panic(err)
	}

	return todos
}

func saveTodos(todos []todo) {
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./todos.json", todoBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func delTodo(todos []todo, id int) []todo {
	for index, todo := range todos {
		if id == todo.Id {
			todos = append(todos[:index], todos[1+index:]...)
		}

	}

	return todos
}

func upddated(todos []todo, id int, completed string, task string) []todo {

	for index, todo := range todos {
		if id == todo.Id {
			todos = append(todos[:index], todos[1+index:]...)
			if id == todo.Id && task != "" && completed != "" {
				todo.Task = task
				todo.Completed = completed
				todos = append(todos, todo)
			} else if id == todo.Id && task != "" {
				todo.Task = task
				todos = append(todos, todo)
			} else if id == todo.Id && completed != "" {
				todo.Completed = completed
				todos = append(todos, todo)
			}

		}

	}

	return todos
}
