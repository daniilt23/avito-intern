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

func NewService(iTeamRepo repository.ITeamRepo,
	iUserRepo repository.IUserRepo,
	iPullRequestRepo repository.IPullRequestRepo,
	logger *slog.Logger,
) *Service {
	return &Service{
		ITeamRepo:    iTeamRepo,
		IUserRepo:    iUserRepo,
		IPullRequest: iPullRequestRepo,
		Logger:       logger,
	}
}
