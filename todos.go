package main

import (
	"encoding/json"
	"io/ioutil"
)

type todo struct {
	Id        string
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
