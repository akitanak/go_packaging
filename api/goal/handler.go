package goal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/akitanak/go_packaging/usecase/goal"
)

type GoalOperationHandler struct {
	usecase *goal.GoalOperations
}

func NewGoalOperationHandler(usecase *goal.GoalOperations) *GoalOperationHandler {
	return &GoalOperationHandler{
		usecase: usecase,
	}
}

func (g *GoalOperationHandler) CreateGoal(ctx *gin.Context) {
	var createGoalReq CreateGoalRequest
	if err := ctx.ShouldBindJSON(&createGoalReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goal, err := g.usecase.CreateGoal(ctx, userID, createGoalReq.Name, createGoalReq.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, goalToGinHash(goal))
}

func (g *GoalOperationHandler) ListGoals(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goals, err := g.usecase.ListGoals(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var res []gin.H
	for _, g := range goals {
		res = append(res, goalToGinHash(g))
	}
	ctx.JSON(http.StatusOK, res)
}

func (g *GoalOperationHandler) GetGoal(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goalID, err := uuid.Parse(ctx.Param("goalID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goal, err := g.usecase.GetGoal(ctx, userID, goalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, goalToGinHash(goal))
}

func (g *GoalOperationHandler) AddTask(ctx *gin.Context) {
	var addTaskReq AddTaskRequest
	if err := ctx.ShouldBindJSON(&addTaskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	goalID, err := uuid.Parse(ctx.Param("goalID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goal, err := g.usecase.AddTask(ctx, userID, goalID, addTaskReq.Name, addTaskReq.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, goalToGinHash(goal))

}
