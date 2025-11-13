package service

import (
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"avito/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func (s *Service) CreateTeam(req *dto.CreateTeamDto) error {
	teamModel := &models.TeamModel{
		TeamName: req.TeamName,
	}

	existsName, err := s.ITeamRepo.GetTeamByName(req.TeamName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if existsName == req.TeamName {
		return fmt.Errorf("%w", apperrors.TeamExists)
	}

	var userModels []models.UserModel
	for _, user := range req.Members {
		createLogger := s.Logger.With("route", "/team/add")
		createLogger.Debug("user data", "user value", user)
		userModels = append(userModels, models.UserModel{
			UserId:   user.UserId,
			Username: user.Username,
			IsActive: user.IsActive,
			TeamName: req.TeamName,
		})
	}

	err = s.ITeamRepo.CreateTeam(teamModel, userModels)
	if err != nil {
		return err
	}

	return nil
}
