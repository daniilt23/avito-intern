package service

import (
	"avito/internal/repository"
	"log/slog"
)

type Service struct {
	IUserRepo    repository.IUserRepo
	ITeamRepo    repository.ITeamRepo
	IPullRequest repository.IPullRequestRepo
	Logger       *slog.Logger
}

func NewService(ITeamRepo repository.ITeamRepo,
	IUserRepo repository.IUserRepo,
	IPullRequestRepo repository.IPullRequestRepo,
	Logger *slog.Logger,
) *Service {
	return &Service{
		ITeamRepo:    ITeamRepo,
		IUserRepo:    IUserRepo,
		IPullRequest: IPullRequestRepo,
		Logger:       Logger,
	}
}
