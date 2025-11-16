package dto

// @Description Create user request
type UserCreateReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

// @Description User response
type UserResponse struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	TeamName string `json:"team_name,omitempty"`
}

// @Description User set active request
type UserSetActiveReq struct {
	UserId   string `json:"user_id" binding:"required"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

// @Description Get review response
type GetReviewResponse struct {
	UserId       string                 `json:"user_id"`
	PullRequests []PullRequestRespShort `json:"pull_requests"`
}
