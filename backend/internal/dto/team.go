package dto

type CreateTeamDto struct {
	TeamName string          `json:"team_name" binding:"required"`
	Members  []CreateUserDto `json:"members" binding:"required,dive"`
}

type CreateUserDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
