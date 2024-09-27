package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func ValidateStatus(status string) error {
	switch status {
	case "done", "todo", "in-progress":
		return nil
	default:
		return fmt.Errorf("Invalid status: %s\nValid statuses: done, todo, in-progress", status)
	}
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CheckTasksFile() error {
	if _, err := os.Stat("tasks.json"); err != nil {
		fmt.Println("tasks.json file not found")
		return CreateTasksFile()
	}
	return nil
}

func CreateTasksFile() error {
	fmt.Println("Creating tasks.json file...")
	if _, err := os.Create("tasks.json"); err != nil {
		return fmt.Errorf("Error creating tasks.json file: %s", err.Error())
	}
	return nil
}

func WriteTasksFile(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %s", err.Error())
	}
	if err := os.WriteFile("tasks.json", data, 0644); err != nil {
		return fmt.Errorf("Error writing tasks.json file: %s", err.Error())
	}
	return nil
}
