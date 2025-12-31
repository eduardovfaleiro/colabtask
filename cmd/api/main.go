package main

import (
	"database/sql"
	"os"

	"github.com/eduardovfaleiro/colabtask/internal/handler"
	"github.com/eduardovfaleiro/colabtask/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
	}
	taskRepo := repository.NewPostgreTaskRepository(db)
	taskHandler := handler.NewTaskHandler(taskRepo)

	projectRepo := repository.NewPostgreProjectRepository(db)
	projectHandler := handler.NewProjectHandler(projectRepo)

	r := gin.Default()

	r.POST("/tasks", taskHandler.Create)
	r.POST("/projects", projectHandler.Create)
	r.Run(":8080")
}
