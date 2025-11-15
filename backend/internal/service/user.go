package service

import (
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"avito/internal/models"
	"database/sql"
	"errors"
)

func (s *Service) GetUsersByTeam(name string) ([]dto.UserResponse, error) {
	s.Logger.Info("Starting function: GetUsersByTeam",
		"team_name", name)
	users, err := s.IUserRepo.GetUsersByTeam(name)
	if err != nil {
		s.Logger.Error("Failed to get users",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Get users")

	var usersDto []dto.UserResponse
	for _, user := range *users {
		usersDto = append(usersDto, *s.UserFromModelToDto(&user))
	}

	s.Logger.Info("Successfully get users by team")

	return usersDto, nil
}

func (s *Service) SetIsActive(req dto.UserSetActiveReq) (*dto.UserResponse, error) {
	s.Logger.Info("Starting function: SetIsActive",
		"user_id", req.UserId)
	user, err := s.IUserRepo.GetUserById(req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("User not found",
				"user_id", req.UserId)
			return nil, apperrors.ErrUserNotFound
		}
		s.Logger.Error("Failed to get user",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Get user")

	err = s.IUserRepo.ChangeActiveStatusById(req.UserId, *req.IsActive)
	if err != nil {
		s.Logger.Error("Failed to change activity",
			"user_id", req.UserId)
		return nil, err
	}

	user.IsActive = *req.IsActive

	userDto := s.UserFromModelToDto(&user)

	s.Logger.Info("Successfully change activity to user",
		"user_id", req.UserId)

	return userDto, nil
}

func (s *Service) UserFromModelToDto(model *models.UserModel) *dto.UserResponse {
	return &dto.UserResponse{
		UserId:   model.UserId,
		Username: model.Username,
		IsActive: model.IsActive,
		TeamName: model.TeamName,
	}
}
