package goal

import (
	"github.com/akitanak/go_packaging/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateGoalRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type AddTaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type Goal struct {
	GoalID      string `uri:"goalID" binding:"required,uuid"`
	UserID      string `uri:"userID" binding:"required,uuid"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
}

type Task struct {
	TaskID      string `json:"taskID" binding:"required,uuid"`
	UserID      string `json:"userID" binding:"required,uuid"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required"`
}

func goalToGinHash(goal *entities.Goal) gin.H {
	var tasks []gin.H
	for _, t := range goal.Tasks {
		tasks = append(tasks, taskToGinHash(t))
	}
	return gin.H{
		"goalID":      goal.GoalID,
		"userID":      goal.UserID,
		"name":        goal.Name,
		"description": goal.Description,
		"tasks":       tasks,
	}
}

func taskToGinHash(task *entities.Task) gin.H {
	return gin.H{
		"taskID":      task.TaskID,
		"userID":      task.UserID,
		"name":        task.Name,
		"description": task.Description,
		"status":      task.Status,
	}
}
