package repositories

import (
	"context"

	"github.com/akitanak/go_packaging/domain/entities"
	"github.com/google/uuid"
)

type GoalRepository interface {
	Get(ctx context.Context, userID uuid.UUID, goalID uuid.UUID) (*entities.Goal, error)
	List(ctx context.Context, userID uuid.UUID) ([]*entities.Goal, error)
	Store(ctx context.Context, goal *entities.Goal) error
	Update(ctx context.Context, goal *entities.Goal) error
}
