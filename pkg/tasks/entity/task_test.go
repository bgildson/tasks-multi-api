package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTask(t *testing.T) {
	description := "my task"
	task, err := NewTask(description)
	if err != nil {
		t.Errorf("an unexpected error occured: %v", err)
	}

	t.Run("creates properly with a valid uuid v4", func(t *testing.T) {
		if _, err := uuid.Parse(task.ID); err != nil {
			t.Errorf("was expecting create the task with a valid uuid v4, but found \"%s\"", task.ID)
		}
	})

	t.Run("creates a task with passed description", func(t *testing.T) {
		if task.Description != description {
			t.Errorf("was expecting \"%s\" description, but found \"%s\"", description, task.Description)
		}
	})

	t.Run("creates a task with not done", func(t *testing.T) {
		if task.Done {
			t.Errorf("the created task should not be done")
		}
	})
}
