package handler

import (
	"net/http"

	"github.com/eduardovfaleiro/colabtask/internal/model"
	"github.com/eduardovfaleiro/colabtask/internal/repository"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo}
}

func (h *TaskHandler) Create(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(c.Request.Context(), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar tarefa"})
		return
	}

	c.JSON(http.StatusCreated, task)
}
