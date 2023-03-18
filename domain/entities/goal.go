package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type GoalStatus string

var (
	GoalStatusSet        GoalStatus = "set"
	GoalStatusInProgress GoalStatus = "in_progress"
	GoalStatusAchieved   GoalStatus = "achieved"
	GoalStatusArchived   GoalStatus = "archived"
)

type Goal struct {
	GoalID      uuid.UUID
	UserID      uuid.UUID
	Name        string
	Description string
	Tasks       []*Task
	Status      GoalStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewGoal(userID uuid.UUID, name, description string) (*Goal, error) {
	if err := validateNewGoal(name, description); err != nil {
		return nil, fmt.Errorf("failed to validate new Goal: %w", err)
	}

	now := time.Now().UTC()
	return &Goal{
		GoalID:      uuid.New(),
		UserID:      userID,
		Name:        name,
		Description: description,
		Tasks:       []*Task{},
		Status:      GoalStatusSet,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func validateNewGoal(name, description string) error {
	return nil
}

func (g *Goal) AddTask(name, descriptiton string) error {
	t, err := NewTask(g.UserID, g.GoalID, name, descriptiton)
	if err != nil {
		return fmt.Errorf("failed to NewTask: %w", err)
	}

	g.Tasks = append(g.Tasks, t)
	g.UpdatedAt = time.Now().UTC()
	return nil
}

func (g *Goal) Achieved() error {
	for _, t := range g.Tasks {
		if !t.IsDone() {
			return fmt.Errorf("all tasks should be done")
		}
	}

	g.Status = GoalStatusAchieved
	g.UpdatedAt = time.Now().UTC()
	return nil
}

func (g *Goal) IsAchieved() bool {
	return g.Status == GoalStatusAchieved
}

func (g *Goal) Archived() {
	g.Status = GoalStatusArchived
	g.UpdatedAt = time.Now().UTC()
}
