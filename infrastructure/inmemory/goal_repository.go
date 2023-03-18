package inmemory

import (
	"context"
	"fmt"

	"github.com/akitanak/go_packaging/domain/entities"
	"github.com/akitanak/go_packaging/domain/repositories"
	"github.com/google/uuid"
)

type InMemoryGoalRepository struct {
	goalsByUserID map[uuid.UUID]map[uuid.UUID]*entities.Goal
}

func NewInMemoryRepository() repositories.GoalRepository {
	return &InMemoryGoalRepository{
		goalsByUserID: make(map[uuid.UUID]map[uuid.UUID]*entities.Goal),
	}
}

func (r *InMemoryGoalRepository) Get(ctx context.Context, userID uuid.UUID, goalID uuid.UUID) (*entities.Goal, error) {
	goals, ok := r.goalsByUserID[userID]
	if !ok {
		return nil, fmt.Errorf("user was not found. userID: %s, goalID: %s", userID, goalID)
	}

	goal, ok := goals[goalID]
	if !ok {
		return nil, fmt.Errorf("goal was not found. userID: %s, goalID: %s", userID, goalID)
	}

	return goal, nil
}

func (r *InMemoryGoalRepository) List(ctx context.Context, userID uuid.UUID) ([]*entities.Goal, error) {
	goalByGoalID, ok := r.goalsByUserID[userID]
	if !ok {
		return nil, fmt.Errorf("user was not found. userID: %s", userID)
	}

	goals := []*entities.Goal{}
	for _, g := range goalByGoalID {
		goals = append(goals, g)
	}
	return goals, nil
}

func (r *InMemoryGoalRepository) Store(ctx context.Context, goal *entities.Goal) error {
	goalByGoalID, ok := r.goalsByUserID[goal.UserID]
	if !ok {
		goalByGoalID = make(map[uuid.UUID]*entities.Goal)
		r.goalsByUserID[goal.UserID] = goalByGoalID
	}
	_, ok = goalByGoalID[goal.GoalID]
	if ok {
		return fmt.Errorf("already stored GoalID: %s", goal.GoalID)
	}
	goalByGoalID[goal.GoalID] = goal
	r.goalsByUserID[goal.UserID] = goalByGoalID
	return nil
}

func (r *InMemoryGoalRepository) Update(ctx context.Context, goal *entities.Goal) error {
	goalByGoalID, ok := r.goalsByUserID[goal.UserID]
	if !ok {
		return fmt.Errorf("goals have not created yet. UserID: %s", goal.UserID)
	}
	_, ok = goalByGoalID[goal.GoalID]
	if !ok {
		return fmt.Errorf("the goal has not created yet. GoalID: %s", goal.GoalID)
	}
	r.goalsByUserID[goal.UserID][goal.GoalID] = goal
	return nil
}
