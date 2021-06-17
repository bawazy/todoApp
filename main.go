package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// todo get subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all Todos")
	getID := getCmd.Int("id", 0, "Todo ID")
	// todo add subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	//input for todo add command
	addId := addCmd.Int("id", 0, "Todo ID")
	addTask := addCmd.String("task", "", "Todo Task (desc)")
	addComp := addCmd.String("completed", "", "Todo Completion")

	//delete subcommand
	delCmd := flag.NewFlagSet("del", flag.ExitOnError)
	// input for delete command
	delID := delCmd.Int("id", 0, "Todo ID")

	// Validation to read what was passed and to make sure that the accurate argument has been passed into our application
	if len(os.Args) < 2 {
		fmt.Println("Expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	// checker to see what value our 2nd argument is
	switch os.Args[1] {
	case "get":
		// if it is the 'get' command we call the handleGet here
		HandleGet(getCmd, getAll, getID)
	case "add":
		// if it is the 'add' command, we call the handleAdd here
		HandleAdd(addCmd, addId, addTask, addComp)
	case "del":
		//if it is the 'delete' command we call the handleDel  function here
		handleDel(delCmd, delID)
	default: // if we dont understand the input
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *int) {
	getCmd.Parse(os.Args[2:])
	if !*all && *id == 0 {
		fmt.Print("id is required or specify --all for all todos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	if *all {
		//return all todos
		todos := getTodos()
		fmt.Printf("ID \t Task \t Completed \n")
		for _, todo := range todos {
			fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Task, todo.Completed)
		}
		return
	}

	if *id != 0 {
		todos := getTodos()
		id := *id

		for _, todo := range todos {
			if id == todo.Id {
				fmt.Printf("ID \t Task \t Completed \n")
				fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Task, todo.Completed)
			}
		}

	}

}

func ValidateTodo(addCmd *flag.FlagSet, id *int, task *string, completed *string) {
	addCmd.Parse(os.Args[2:])
	if *id == 0 || *task == "" || *completed == "" {
		fmt.Print("all fields are required for adding a todo \n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateDel(addCmd *flag.FlagSet, id *int) {
	addCmd.Parse(os.Args[2:])
	if *id == 0 {
		fmt.Print("A valid ID is required a delete a todo \n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}
func handleDel(delCmd *flag.FlagSet, id *int) {
	ValidateDel(delCmd, id)

	if *id != 0 {
		todos := getTodos()
		id := *id
		// fmt.Println(id)
		delTodo(todos, id)
		saveTodos(todos)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *int, task *string, completed *string) {

	ValidateTodo(addCmd, id, task, completed)

	todo := todo{
		Id:        *id,
		Task:      *task,
		Completed: *completed,
	}
	todos := getTodos()

	todos = append(todos, todo)
	saveTodos(todos)

}
