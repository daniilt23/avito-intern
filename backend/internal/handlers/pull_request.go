package handlers

import (
	"avito/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary      Create PR
//	@Description  For use this method you need to be admin
//	@Tags         Pull request
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.CreatePullRequestReq  true  "request body with all required fields"
//	@Success      201  {array}   dto.CreatePullRequestResp
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /pullRequest/create [post]
//  @Security 	  BearerAuth
func (h *Handler) CreatePullRequest(c *gin.Context) {
	var req dto.CreatePullRequestReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	reviewersId, err := h.srv.CreatePullRequest(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pr": dto.CreatePullRequestResp{
		PullRequestId:   req.PullRequestId,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorId,
		Status:          "OPEN",
		Reviewers:       reviewersId,
	}})
}

//	@Summary      Set PR status to MERGED
//	@Description  For use this method you need to be admin
//	@Tags         Pull request
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.SetMergeStatusReq  true  "request body with all required fields"
//	@Success      200  {array}   dto.PullRequestResp
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /pullRequest/merge [patch]
//  @Security 	  BearerAuth
func (h *Handler) SetMergeStatus(c *gin.Context) {
	var req dto.SetMergeStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	pullRequest, err := h.srv.SetMergeStatus(req)
	if err != nil {
		h.handleError(c, err)
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

//  @Summary      Reassigne reviewer
//	@Description  For use this method you need to be admin
//	@Tags         Pull request
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.ReassignReviewerReq  true  "request body with all required fields"
//	@Success      200  {array}   dto.ReassignReviewerResp
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      409  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /pullRequest/reassign [post]
//  @Security 	  BearerAuth
func (h *Handler) ReassignReviewer(c *gin.Context) {
	var req dto.ReassignReviewerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	resp, err := h.srv.ReassignReviewer(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pr":          resp,
		"replaced_by": req.UserId})
}
