package goal

import (
	"context"
	"fmt"

	"github.com/akitanak/go_packaging/domain/entities"
	"github.com/akitanak/go_packaging/domain/repositories"
	"github.com/google/uuid"
)

type GoalOperations struct {
	repo repositories.GoalRepository
}

func NewGoalOperations(repo repositories.GoalRepository) *GoalOperations {
	return &GoalOperations{
		repo: repo,
	}
}

func (g *GoalOperations) CreateGoal(ctx context.Context, userID uuid.UUID, name, description string) (*entities.Goal, error) {
	goal, err := entities.NewGoal(userID, name, description)
	if err != nil {
		return nil, fmt.Errorf("failed to NewGoal: %w", err)
	}

	if err := g.repo.Store(ctx, goal); err != nil {
		return nil, fmt.Errorf("failed to store goal: %w", err)
	}

	return goal, nil
}

func (g *GoalOperations) ListGoals(ctx context.Context, userID uuid.UUID) ([]*entities.Goal, error) {
	goals, err := g.repo.List(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list goals: %w", err)
	}

	return goals, nil
}

func (g *GoalOperations) GetGoal(ctx context.Context, userID, goalID uuid.UUID) (*entities.Goal, error) {
	goal, err := g.repo.Get(ctx, userID, goalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get goal: %w", err)
	}

	return goal, nil
}

func (g *GoalOperations) AddTask(ctx context.Context, userID, goalID uuid.UUID, taskName, taskDescription string) (*entities.Goal, error) {
	goal, err := g.repo.Get(ctx, userID, goalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get goal: %w", err)
	}

	if err := goal.AddTask(taskName, taskDescription); err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}

	if err := g.repo.Update(ctx, goal); err != nil {
		return nil, fmt.Errorf("failed to store goal: %w", err)
	}

	return goal, nil
}
