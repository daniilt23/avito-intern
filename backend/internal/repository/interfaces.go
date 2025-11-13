package repository

import "avito/internal/models"

type ITeamRepo interface {
	CreateTeam(reqTeam *models.TeamModel, reqUsers []models.UserModel) error
	GetTeamByName(name string) (string, error)
}

type IUserRepo interface {
	CreateUser(name string, isActive bool, teamId int) error
}
