package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(isAdminOnly bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			respondUnauthorized(c, apperrors.ErrMissingHeader)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			respondUnauthorized(c, apperrors.ErrInvalidFormat)
			return
		}

		token := headerParts[1]
		if isAdminOnly {
			if token != os.Getenv("ADMIN_TOKEN") {
				respondUnauthorized(c, apperrors.ErrUnauthorized)
				return
			}
		} else {
			if token != os.Getenv("USER_TOKEN") && token != os.Getenv("ADMIN_TOKEN") {
				respondUnauthorized(c, apperrors.ErrUnauthorized)
				return
			}
		}

		c.Next()
	}
}

func respondUnauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": dto.ErrorResponse{
			Code:    consts.ErrUnauthorized,
			Message: err.Error(),
		},
	})
}
