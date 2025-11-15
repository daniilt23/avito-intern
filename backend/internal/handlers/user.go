package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetIsActive(c *gin.Context) {
	var req dto.UserSetActiveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: err.Error(),
		}})
		return
	}

	user, err := h.srv.SetIsActive(req)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrUserNotFound.Error(),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": dto.UserResponse{
		UserId: user.UserId,
		Username: user.Username,
		IsActive: user.IsActive,
		TeamName: user.TeamName,
	}})
}
