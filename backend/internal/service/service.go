package service

import (
	"avito/internal/repository"
	"log/slog"
)

type Service struct {
	IUserRepo repository.IUserRepo
	ITeamRepo repository.ITeamRepo
	Logger    *slog.Logger
}

func NewService(ITeamRepo repository.ITeamRepo, IUserRepo repository.IUserRepo, Logger *slog.Logger) *Service {
	return &Service{
		ITeamRepo: ITeamRepo,
		IUserRepo: IUserRepo,
		Logger: Logger,
	}
}
