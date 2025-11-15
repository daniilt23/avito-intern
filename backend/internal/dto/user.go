package dto

type UserCreateReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

type UserResponse struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	TeamName string `json:"team_name,omitempty"`
}

type UserSetActiveReq struct {
	UserId   string `json:"user_id" binding:"required"`
	IsActive *bool  `json:"is_active" binding:"required"`
}
