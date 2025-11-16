package dto

import "time"

// @Description Create PR request
type CreatePullRequestReq struct {
	PullRequestId   string `json:"pull_request_id" binding:"required"`
	PullRequestName string `json:"pull_request_name" binding:"required"`
	AuthorId        string `json:"author_id" binding:"required"`
}

// @Description PR response
type CreatePullRequestResp struct {
	PullRequestId   string    `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorId        string    `json:"author_id"`
	Status          string    `json:"status"`
	Reviewers       []string  `json:"assigned_reviewers,omitempty"`
}

// @Description PR response
type PullRequestResp struct {
	PullRequestId   string    `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorId        string    `json:"author_id"`
	Status          string    `json:"status"`
	Reviewers       []string  `json:"assigned_reviewers,omitempty"`
	MergedAt        time.Time `json:"merged_at,omitzero"`
}

// @Description PR response short
type PullRequestRespShort struct {
	PullRequestId   string    `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorId        string    `json:"author_id"`
	Status          string    `json:"status"`
}

// @Description Set merge status request
type SetMergeStatusReq struct {
	PullRequestId string `json:"pull_request_id" binding:"required"`
}

// @Description Reassigne reviewer request
type ReassignReviewerReq struct {
	PullRequestId string `json:"pull_request_id" binding:"required"`
	UserId        string `json:"user_id" binding:"required"`
}

// @Description Reassigne reviewer response
type ReassignReviewerResp struct {
	PullRequestId   string    `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorId        string    `json:"author_id"`
	Status          string    `json:"status"`
	Reviewers       []string  `json:"assigned_reviewers,omitempty"`
	ReplacedBy      string    `json:"replaced_by"`
}
