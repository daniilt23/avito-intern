package dto

// @Description Error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// @Description Team request
type TeamReq struct {
	TeamName string          `json:"team_name" binding:"required"`
	Members  []UserCreateReq `json:"members" binding:"required,dive"`
}

// @Description Team response
type TeamResp struct {
	TeamName string         `json:"team_name"`
	Members  []UserResponse `json:"members"`
}
