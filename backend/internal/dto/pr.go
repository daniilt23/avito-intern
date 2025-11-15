package dto

import "time"

type CreatePullRequestReq struct {
	PullRequestId   string `json:"pull_request_id" binding:"required"`
	PullRequestName string `json:"pull_request_name" binding:"required"`
	AuthorId        string `json:"author_id" binding:"required"`
}

type PullRequestResp struct {
	PullRequestId   string    `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorId        string    `json:"author_id"`
	Status          string    `json:"status"`
	Reviewers       []string  `json:"assigned_reviewers,omitempty"`
	MergedAt        time.Time `json:"merged_at,omitzero"`
}

type SetMergeStatusReq struct {
	PullRequestId string `json:"pull_request_id" binding:"required"`
}

type ReassignReviewerReq struct {
	PullRequestId string `json:"pull_request_id" binding:"required"`
	UserId        string `json:"user_id" binding:"required"`
}
