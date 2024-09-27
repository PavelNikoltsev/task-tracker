package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	deleteTasksFile()
	code := m.Run()
	deleteTasksFile()
	os.Exit(code)
}

func deleteTasksFile() {
	_ = os.Remove("../tasks.json")
	_ = os.Remove("tasks.json")
}

func TestCheckTasksFile(t *testing.T) {
	deleteTasksFile()
	err := CreateTasksFile()
	if err != nil {
		t.Errorf("CreateTasksFile() error = %v", err)
	}
	err = CheckTasksFile()
	if err != nil {
		t.Errorf("CheckTasksFile() error = %v", err)
	}
}

func TestCreateTasksFile(t *testing.T) {
	deleteTasksFile()
	err := CreateTasksFile()
	if err != nil {
		t.Errorf("CreateTasksFile() error = %v", err)
	}
	if _, err := os.Stat("tasks.json"); err != nil {
		t.Errorf("tasks.json file not found")
	}
}

func TestValidateStatus(t *testing.T) {
	deleteTasksFile()
	err := ValidateStatus("done")
	if err != nil {
		t.Errorf("ValidateStatus() error = %v", err)
	}
	err = ValidateStatus("invalid")
	if err == nil {
		t.Errorf("ValidateStatus() error = %v", err)
	}
}
