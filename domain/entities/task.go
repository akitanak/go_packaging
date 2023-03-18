package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

var (
	TaskStatusToDo     TaskStatus = "todo"
	TaskStatusDone     TaskStatus = "done"
	TaskStatusArchived TaskStatus = "archived"
)

type Task struct {
	TaskID      uuid.UUID
	UserID      uuid.UUID
	GoalID      uuid.UUID
	Name        string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(userID, goalID uuid.UUID, name, description string) (*Task, error) {
	if err := validateNewTask(name, description); err != nil {
		return nil, fmt.Errorf("failed to validate new Task: %w", err)
	}

	now := time.Now().UTC()
	return &Task{
		TaskID:      uuid.New(),
		UserID:      userID,
		GoalID:      goalID,
		Name:        name,
		Description: description,
		Status:      TaskStatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func validateNewTask(name, description string) error {
	return nil
}

func (t *Task) Done() error {
	if t.Status != TaskStatusToDo {
		return fmt.Errorf("can be Done only from the ToDo state")
	}

	t.Status = TaskStatusDone
	t.UpdatedAt = time.Now().UTC()
	return nil
}

func (t *Task) IsDone() bool {
	return t.Status == TaskStatusDone
}

func (t *Task) Archived() {
	t.Status = TaskStatusArchived
	t.UpdatedAt = time.Now().UTC()
}

func (t *Task) IsArchived() bool {
	return t.Status == TaskStatusArchived
}
