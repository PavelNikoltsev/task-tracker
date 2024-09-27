package commands

import (
	"errors"
	"fmt"
	"os"
	"task-tracker/models"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	deleteTasksFile()
	os.Exit(code)
}

func setup() {
	deleteTasksFile()
	_ = models.CheckTasksFile()
}

func deleteTasksFile() {
	fmt.Println(1)
	_ = os.Remove("../tasks.json")
	_ = os.Remove("tasks.json")
	fmt.Println(2)
}

func TestGetAllTasksEmptyFile(t *testing.T) {
	deleteTasksFile()
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 0 {
		t.Error(errors.New("tasks should be empty"))
	}
}

func TestGetAllTasksWithRecords(t *testing.T) {
	deleteTasksFile()
	tasks := []models.Task{
		{
			ID:          1,
			Description: t.Name(),
		},
	}
	if err := models.WriteTasksFile(tasks); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 1 {
		t.Error(errors.New("tasks should be 1"))
	}
}

func TestGetTask(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		ID:          1,
		Description: t.Name(),
	}
	if err := models.WriteTasksFile([]models.Task{task}); err != nil {
		t.Error(err)
	}
	task, err := GetTask(1)
	if err != nil {
		t.Error(err)
	}
	if task.ID != 1 {
		t.Error(errors.New("task ID should be 1"))
	}
	if task.Description != t.Name() {
		t.Error(errors.New("task description should be " + t.Name()))
	}
}

func TestAdd(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		Description: t.Name(),
	}
	if err := Add(task.Description); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if tasks[0].ID != 1 {
		t.Error(errors.New("task ID should be 1"))
	}
	if tasks[0].Description != t.Name() {
		fmt.Println(123, tasks[0])
		t.Error(errors.New("task description should be " + t.Name()))
	}
}

func TestUpdate(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		ID:          1,
		Description: t.Name(),
	}
	if err := models.WriteTasksFile([]models.Task{task}); err != nil {
		t.Error(err)
	}
	if err := Update(1, "new description"); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if tasks[0].Description != "new description" {
		t.Error(errors.New("task description should be new description"))
	}
}

func TestDelete(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		ID:          1,
		Description: t.Name(),
	}
	if err := models.WriteTasksFile([]models.Task{task}); err != nil {
		t.Error(err)
	}
	if err := Delete(1); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 0 {
		t.Error(errors.New("tasks should be 0"))
	}
}

func TestClear(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		ID:          1,
		Description: t.Name(),
	}
	if err := models.WriteTasksFile([]models.Task{task}); err != nil {
		t.Error(err)
	}
	if err := Clear(); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 0 {
		t.Error(errors.New("tasks should be 0"))
	}
}

func TestSetStatus(t *testing.T) {
	deleteTasksFile()
	task := models.Task{
		ID:          1,
		Description: t.Name(),
	}
	if err := models.WriteTasksFile([]models.Task{task}); err != nil {
		t.Error(err)
	}
	if err := SetStatus(1, "done"); err != nil {
		t.Error(err)
	}
	tasks, err := GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	if tasks[0].Status != models.StatusDone {
		t.Error(errors.New("task status should be done"))
	}
}
