package service

import (
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"avito/internal/models"
	"database/sql"
	"errors"
)

func (s *Service) CreateTeam(req *dto.TeamReq) error {
	s.Logger.Info("Starting function: CreateTeam",
		"team_name", req.TeamName)
	teamModel := &models.TeamModel{
		TeamName: req.TeamName,
	}

	existsName, err := s.ITeamRepo.GetTeamByName(req.TeamName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.Logger.Error("Failed to get team",
			"error", err)
		return err
	}

	if existsName == req.TeamName {
		s.Logger.Warn("Team exists",
			"team_name", req.TeamName)
		return apperrors.ErrTeamExists
	}

	var userModels []models.UserModel
	for _, user := range req.Members {
		s.Logger.Debug("User data", "user value", user)
		userModels = append(userModels, models.UserModel{
			UserId:   user.UserId,
			Username: user.Username,
			IsActive: *user.IsActive,
			TeamName: req.TeamName,
		})
	}

	err = s.ITeamRepo.CreateTeam(teamModel, userModels)
	if err != nil {
		s.Logger.Error("Failed to create team",
			"error", err)
		return err
	}

	s.Logger.Info("Successfully create team",
		"team_name", req.TeamName)

	return nil
}

func (s *Service) GetTeamByName(name string) ([]dto.UserResponse, error) {
	s.Logger.Info("Starting function: GetTeamByName",
		"name", name)
	if _, err := s.ITeamRepo.GetTeamByName(name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("Team not found",
				"team_name", name)
			return nil, apperrors.ErrTeamNotFound
		}
		s.Logger.Error("Failed to get team",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Get team")

	users, err := s.GetUsersByTeam(name)
	if err != nil {
		s.Logger.Error("Failed to get users by team",
			"error", err)
		return nil, err
	}

	s.Logger.Info("Successfully get team by name")

	return users, nil
}
