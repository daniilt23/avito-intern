package dto

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type TeamReq struct {
	TeamName string          `json:"team_name" binding:"required"`
	Members  []UserCreateReq `json:"members" binding:"required,dive"`
}

type TeamResp struct {
	TeamName string         `json:"team_name"`
	Members  []UserResponse `json:"members"`
}
