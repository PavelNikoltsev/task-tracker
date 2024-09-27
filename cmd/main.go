package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/commands"
	"task-tracker/models"
)

func getID(stringID string) (int, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return 0, fmt.Errorf("Task ID should be a valid number: %v\n", err)
	}
	return id, nil
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task description> - Add a new task")
	fmt.Println("  update <task ID> <task description> - Update a task")
	fmt.Println("  delete <task ID> - Delete a task")
	fmt.Println("  clear - Clear all tasks")
	fmt.Println("  list [status] - List tasks")
	fmt.Println("  show <task ID> - Show details of a task")
	fmt.Println("  mark-done <task ID> - Mark a task as done")
	fmt.Println("  mark-todo <task ID> - Mark a task as todo")
	fmt.Println("  mark-in-progress <task ID> - Mark a task as in progress")
	fmt.Println("  help - Show this help message")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("It is necessary to transmit a command.")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Task description is required for add command.")
			return
		}
		taskDesc := os.Args[2]
		err := commands.Add(taskDesc)
		if err != nil {
			fmt.Println(err)
			break
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Task ID and description are required for update command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		taskDesc := os.Args[3]
		err = commands.Update(taskID, taskDesc)
		if err != nil {
			fmt.Println(err)
			break
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required for delete command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		err = commands.Delete(taskID)
		if err != nil {
			fmt.Println(err)
			break
		}
	case "clear":
		err := commands.Clear()
		if err != nil {
			fmt.Println(err)
			break
		}
	case "list":
		statusArg := ""
		if len(os.Args) > 2 {
			statusArg = os.Args[2]
		}
		err := commands.List(statusArg)
		if err != nil {
			fmt.Println(err)
			break
		}
	case "show":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required for show command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		err = commands.Show(taskID)
		if err != nil {
			fmt.Println(err)
			break
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required for mark-done command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		err = commands.SetStatus(taskID, string(models.StatusDone))
		if err != nil {
			fmt.Println(err)
			break
		}
	case "mark-todo":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required for mark-todo command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		err = commands.SetStatus(taskID, string(models.StatusTodo))
		if err != nil {
			fmt.Println(err)
			break
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required for mark-in-progress command.")
			return
		}
		taskID, err := getID(os.Args[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		err = commands.SetStatus(taskID, string(models.StatusInProgress))
		if err != nil {
			fmt.Println(err)
			break
		}
	case "help":
		help()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		help()
	}
}
