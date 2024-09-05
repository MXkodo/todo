package main

import (
	"fmt"

	"github.com/MXkodo/todo/config"

	"github.com/MXkodo/todo/internal/handler"
	"github.com/MXkodo/todo/internal/repo"
	"github.com/MXkodo/todo/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("/app/config/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	db, err := repo.NewDBConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to create database connection: %v", err))
	}

	if err := repo.Migrate(db); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	taskRepo := repo.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.Default()
	handler.RegisterRoutes(r, taskHandler)

	r.Run(":" + cfg.Server.Port)
}
