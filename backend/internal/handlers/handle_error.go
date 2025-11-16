package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrPrNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    consts.ErrNotFound,
			Message: apperrors.ErrPrNotFound.Error(),
		})
	case errors.Is(err, apperrors.ErrUserNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    consts.ErrNotFound,
			Message: apperrors.ErrUserNotFound.Error(),
		})
	case errors.Is(err, apperrors.ErrPrExists):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: apperrors.ErrPrExists.Error(),
		})
	case errors.Is(err, apperrors.ErrPrMerged):
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Code:    consts.ErrPrMerged,
			Message: apperrors.ErrPrMerged.Error(),
		})
	case errors.Is(err, apperrors.ErrNotAssigned):
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Code:    consts.ErrNotAssigned,
			Message: apperrors.ErrNotAssigned.Error(),
		})
	case errors.Is(err, apperrors.ErrNoCandidate):
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Code:    consts.ErrNoCandidate,
			Message: apperrors.ErrNoCandidate.Error(),
		})
	case errors.Is(err, apperrors.ErrTeamExists):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    consts.ErrTeamExists,
			Message: apperrors.ErrTeamExists.Error(),
		})
	case errors.Is(err, apperrors.ErrTeamNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    consts.ErrNotFound,
			Message: apperrors.ErrTeamNotFound.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		})
	}
}
