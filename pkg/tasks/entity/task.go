package entity

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a task
type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTask creates a new Task instance
func NewTask(description string) (*Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:          id.String(),
		Description: description,
		Done:        false,
	}, nil
}
