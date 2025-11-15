package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePullRequest(c *gin.Context) {
	var req dto.CreatePullRequestReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: err.Error(),
		}})
		return
	}

	reviewersId, err := h.srv.CreatePullRequest(req)
	if err != nil {
		if errors.Is(err, apperrors.ErrPrExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrBadRequest,
				Message: apperrors.ErrPrExists.Error(),
			}})
			return
		}
		if errors.Is(err, apperrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrUserNotFound.Error(),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: err.Error(),
		}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pr": dto.PullRequestResp{
		PullRequestId:   req.PullRequestId,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorId,
		Status:          "OPEN",
		Reviewers:       reviewersId,
	}})
}

func (h *Handler) SetMergeStatus(c *gin.Context) {
	var req dto.SetMergeStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: err.Error(),
		}})
		return
	}

	pullRequest, err := h.srv.SetMergeStatus(req)
	if err != nil {
		if errors.Is(err, apperrors.ErrPrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrPrNotFound.Error(),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pr": dto.PullRequestResp{
		PullRequestId:   pullRequest.PullRequestId,
		PullRequestName: pullRequest.PullRequestName,
		AuthorId:        pullRequest.AuthorId,
		Status:          pullRequest.Status,
		Reviewers:       pullRequest.Reviewers,
		MergedAt:        pullRequest.MergedAt,
	}})
}

func (h *Handler) ReassignReviewer(c *gin.Context) {
	var req dto.ReassignReviewerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: err.Error(),
		}})
		return
	}

	resp, err := h.srv.ReassignReviewer(req)
	if err != nil {
		if errors.Is(err, apperrors.ErrPrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrPrNotFound.Error(),
			}})
			return
		}
		if errors.Is(err, apperrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotFound,
				Message: apperrors.ErrUserNotFound.Error(),
			}})
			return
		}
		if errors.Is(err, apperrors.ErrPrMerged) {
			c.JSON(http.StatusConflict, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrPrMerged,
				Message: apperrors.ErrPrMerged.Error(),
			}})
			return
		}
		if errors.Is(err, apperrors.ErrNotAssigned) {
			c.JSON(http.StatusConflict, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNotAssigned,
				Message: apperrors.ErrNotAssigned.Error(),
			}})
			return
		}
		if errors.Is(err, apperrors.ErrNoCandidate) {
			c.JSON(http.StatusConflict, gin.H{"error": dto.ErrorResponse{
				Code:    consts.ErrNoCandidate,
				Message: apperrors.ErrNoCandidate.Error(),
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrorResponse{
			Code:    consts.ErrServer,
			Message: apperrors.ErrInternalServer.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pr":          resp,
		"replaced_by": req.UserId})
}

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

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userId,
		"pull_requests": userPrs,
	})
}
