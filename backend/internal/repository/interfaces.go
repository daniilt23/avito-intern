package repository

import (
	"avito/internal/models"
	"time"
)

type ITeamRepo interface {
	CreateTeam(reqTeam *models.TeamModel, reqUsers []models.UserModel) error
	GetTeamByName(name string) (string, error)
}

type IUserRepo interface {
	GetUsersByTeam(name string) (*[]models.UserModel, error)
	GetUserById(id string) (models.UserModel, error)
	ChangeActiveStatusById(id string, isActive bool) error
	GetUsersToPr(teamName string, authorId string) ([]string, error)
	GetUsersByPr(prName string) ([]string, error)
	GetUserToPr(teamName string, authorId string, prId string, oldId string) (string, error)
	ChangeUserAssignee(prId string, userId string, newUserId string) error
}

type IPullRequestRepo interface {
	CreatePullRequest(req *models.PullRequestModel) error
	GetPullRequestById(id string) (*models.PullRequestModel, error)
	AddReviewers(usersId []string, prId string) error
	SetMergeStatus(prId string, mergedAt time.Time) error
	GetPullRequestByUser(userId string, pullRequestId string) (bool, error)
	GetPullRequestsByUser(userId string) ([]models.PullRequestModel, error)
}
