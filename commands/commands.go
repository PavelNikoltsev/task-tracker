package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"task-tracker/models"
	"time"
)

func Add(desc string) error {
	if err := models.CheckTasksFile(); err != nil {
		return err
	}
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	task := models.Task{
		Description: desc,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
	}
	if len(tasks) > 0 {
		task.ID = len(tasks) + 1
	} else {
		task.ID = 1
	}
	tasks = append(tasks, task)
	if err := models.WriteTasksFile(tasks); err != nil {
		return err
	}
	fmt.Println("Task added successfully")
	return nil
}

func Show(id int) error {
	task, err := GetTask(id)
	if err != nil {
		return err
	}
	fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n",
		task.ID, task.Description, task.Status, task.CreatedAt.Format("02/01/2006 15:04:05"), task.UpdatedAt.Format("02/01/2006 15:04:05"))
	return nil
}

func List(status string) error {
	if status == "" {
		if err := printAllTasks(); err != nil {
			return err
		}
	} else {
		if err := printTasksFilteredByStatus(status); err != nil {
			return err
		}
	}
	return nil
}

func printAllTasks() error {
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return errors.New("No tasks found")
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format("02/01/2006 15:04:05"), task.UpdatedAt.Format("02/01/2006 15:04:05"))
	}
	return nil
}

func printTasksFilteredByStatus(status string) error {
	if err := models.ValidateStatus(status); err != nil {
		return err
	}
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return errors.New("No tasks found")
	}
	filteredTasks := []models.Task{}
	for _, task := range tasks {
		if task.Status == models.Status(status) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	if len(filteredTasks) == 0 {
		return errors.New("No tasks found")
	}
	for _, task := range filteredTasks {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format("02/01/2006 15:04:05"), task.UpdatedAt.Format("02/01/2006 15:04:05"))
	}
	return nil
}

func Update(id int, desc string) error {
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	task, err := GetTask(id)
	if err != nil {
		return err
	}
	task.Description = desc
	task.UpdatedAt = time.Now()
	tasks[id-1] = task
	if err = models.WriteTasksFile(tasks); err != nil {
		return err
	}
	fmt.Printf("Task %d updated successfully\n", id)
	return nil
}

func Delete(id int) error {
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return errors.New("Tasks list is empty")
	}
	if id < 1 || id > len(tasks) {
		return errors.New("Invalid Task ID")
	}
	tasks = append(tasks[:id-1], tasks[id:]...)
	if err = models.WriteTasksFile(tasks); err != nil {
		return err
	}
	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

func Clear() error {
	if err := models.CheckTasksFile(); err != nil {
		return err
	}
	if err := models.WriteTasksFile([]models.Task{}); err != nil {
		return err
	}
	fmt.Println("Tasks cleared successfully")
	return nil
}

func SetStatus(id int, status string) error {
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}
	task, err := GetTask(id)
	if err != nil {
		return err
	}
	if err := models.ValidateStatus(status); err != nil {
		return err
	}
	task.Status = models.Status(status)
	task.UpdatedAt = time.Now()
	tasks[id-1] = task
	if err = models.WriteTasksFile(tasks); err != nil {
		return err
	}
	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}

func GetAllTasks() ([]models.Task, error) {
	if err := models.CheckTasksFile(); err != nil {
		return nil, err
	}
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, fmt.Errorf("Error opening tasks.json file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading tasks.json file: %v", err)
	}

	var tasks []models.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		switch err.Error() {
		case "unexpected end of JSON input":
			return tasks, nil
		default:
			return nil, fmt.Errorf("Error unmarshalling tasks.json file: %v", err)
		}
	}
	return tasks, nil
}

func GetTask(id int) (models.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return models.Task{}, err
	}
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task with ID %d not found", id)
}
