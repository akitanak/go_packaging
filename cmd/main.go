package main

import (
	"github.com/gin-gonic/gin"

	"github.com/akitanak/go_packaging/di"
)

func main() {
	r := gin.Default()

	goalHandler := di.InitializeGoalOperations()
	r.POST("/users/:userID/goal", goalHandler.CreateGoal)
	r.GET("/users/:userID/goals", goalHandler.ListGoals)
	r.GET("/users/:userID/goals/:goalID", goalHandler.GetGoal)
	r.POST("/users/:userID/goals/:goalID/task", goalHandler.AddTask)

	r.Run(":8080")
}
