package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary      Change activity to user
//	@Description  Change is_active flag
//	@Tags         User
//	@Accept       json
//	@Produce      json
//  @Param        required    body     dto.UserSetActiveReq  true  "json body to change active"
//	@Success      200  {array}   dto.UserResponse
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /users/setIsActive [patch]
//	@Security 	  BearerAuth
func (h *Handler) SetIsActive(c *gin.Context) {
	var req dto.UserSetActiveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := h.srv.SetIsActive(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": dto.UserResponse{
		UserId:   user.UserId,
		Username: user.Username,
		IsActive: user.IsActive,
		TeamName: user.TeamName,
	}})
}

//	@Summary      Get all pull request by user
//	@Description  For use this method you need to be user/admin
//	@Tags         User
//	@Accept       json
//	@Produce      json
//  @Param        user_id    query     string  true  "query param user_id"
//	@Success      200  {array}   dto.GetReviewResponse
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /users/getReview [get]
//  @Security 	  BearerAuth
func (h *Handler) GetReview(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: apperrors.ErrNoUserQuery.Error(),
		}})
		return
	}

	userPrs, err := h.srv.GetPullRequestsByUser(userId)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userId,
		"pull_requests": userPrs,
	})
}
