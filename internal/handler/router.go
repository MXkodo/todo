package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, taskHandler *TaskHandler) {
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
