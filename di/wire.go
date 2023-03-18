//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	api_goal "github.com/akitanak/go_packaging/api/goal"
	"github.com/akitanak/go_packaging/infrastructure/inmemory"
	uc_goal "github.com/akitanak/go_packaging/usecase/goal"
)

func InitializeGoalOperations() *api_goal.GoalOperationHandler {
	wire.Build(
		api_goal.NewGoalOperationHandler,
		uc_goal.NewGoalOperations,
		inmemory.NewInMemoryRepository,
	)

	return &api_goal.GoalOperationHandler{}
}
