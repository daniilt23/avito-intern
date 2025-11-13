package handlers

import (
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeam(c *gin.Context) {
	var req dto.CreateTeamDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.CreateTeam(&req)
	if err != nil {
		if errors.Is(err, apperrors.TeamExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("team with name: '%s' exists", req.TeamName)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot do operation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"answer": "create team"})
}

func (h *Handler) GetTeam(c *gin.Context) {
	
}
