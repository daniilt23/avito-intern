package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeam(c *gin.Context) {
	var req dto.TeamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.CreateTeam(&req)
	if err != nil {
		if errors.Is(err, apperrors.ErrTeamExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrTeamExists,
				Message: fmt.Sprintf("%s already exists", req.TeamName),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		}})
		return
	}

	users, err := h.srv.GetUsersByTeam(req.TeamName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
			Code: consts.ErrNotFound,
			Message: apperrors.ErrTeamNotFound.Error(),
		}})
	}

	c.JSON(http.StatusCreated, dto.TeamResp{
		TeamName: req.TeamName,
		Members:  users,
	})
}

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: apperrors.ErrNoTeamQuery.Error(),
		}})
		return
	}

	users, err := h.srv.GetTeamByName(teamName)
	if err != nil {
		if errors.Is(err, apperrors.ErrTeamNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrTeamNotFound.Error(),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, dto.TeamResp{
		TeamName: teamName,
		Members:  users,
	})
}
