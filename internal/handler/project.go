package handler

import (
	"net/http"

	"github.com/eduardovfaleiro/colabtask/internal/model"
	"github.com/eduardovfaleiro/colabtask/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	repo repository.ProjectRepository
}

func NewProjectHandler(repo repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{repo}
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(c.Request.Context(), &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar projeto"})
		return
	}

	c.JSON(http.StatusCreated, project)
}
