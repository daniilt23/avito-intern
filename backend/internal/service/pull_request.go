package service

import (
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"avito/internal/models"
	"database/sql"
	"errors"
	"time"
)

func (s *Service) CreatePullRequest(req dto.CreatePullRequestReq) ([]string, error) {
	s.Logger.Info("Start CreatePullRequest function")
	_, err := s.IPullRequest.GetPullRequestById(req.PullRequestId)
	if err == nil {
		s.Logger.Warn("Pull request already exists",
			"pr_id", req.PullRequestId)
		return nil, apperrors.ErrPrExists
	}

	if !errors.Is(err, sql.ErrNoRows) {
		s.Logger.Error("Failed to check PR",
			"pr_id", req.PullRequestId,
			"error", err,
		)
		return nil, err
	}

	author, err := s.IUserRepo.GetUserById(req.AuthorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("Author not found",
				"author_id", req.AuthorId,
			)
			return nil, apperrors.ErrUserNotFound
		}
		s.Logger.Error("Failed to get author",
			"author_id", req.AuthorId,
			"error", err,
		)
		return nil, err
	}

	s.Logger.Debug("Author found",
		"author_id", author.UserId,
		"team", author.TeamName,
		"is_active", author.IsActive,
	)

	model := s.PrFromDtoToModel(&req)
	err = s.IPullRequest.CreatePullRequest(model)
	if err != nil {
		s.Logger.Error("Failed to create pull request",
			"pr_id", req.PullRequestId,
			"error", err,
		)
		return nil, err
	}

	s.Logger.Debug("Pull request created",
		"pr_id", req.PullRequestId,
	)

	usersId, err := s.IUserRepo.GetUsersToPr(author.TeamName, author.UserId)
	if err != nil {
		s.Logger.Error("Failed to get reviewers",
			"team", author.TeamName,
			"author_id", author.UserId,
			"error", err,
		)
		return nil, err
	}

	s.Logger.Info("Found reviewers",
		"count", len(usersId),
		"reviewers", usersId,
		"pr_id", req.PullRequestId,
	)

	err = s.IPullRequest.AddReviewers(usersId, req.PullRequestId)
	if err != nil {
		s.Logger.Error("Failed to add reviewers",
			"pr_id", req.PullRequestId,
			"reviewers", usersId,
			"error", err,
		)
		return nil, err
	}

	s.Logger.Info("Pull request created successfully",
		"pr_id", req.PullRequestId,
		"author_id", req.AuthorId,
		"reviewers_count", len(usersId),
	)

	return usersId, nil
}

func (s *Service) SetMergeStatus(req dto.SetMergeStatusReq) (*dto.PullRequestResp, error) {
	s.Logger.Info("Start function SetMergeStatus",
		"pr_id", req.PullRequestId,
	)
	model, err := s.IPullRequest.GetPullRequestById(req.PullRequestId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("PR not found",
				"pr_id", req.PullRequestId,
				"error", err)
			return nil, apperrors.ErrPrNotFound
		}
		s.Logger.Error("Failed to get PR",
			"pr_id", req.PullRequestId,
			"error", err)
		return nil, err
	}

	s.Logger.Debug("Get PR",
		"PR name", model.PullRequestName)

	users, err := s.IUserRepo.GetUsersByPr(req.PullRequestId)
	if err != nil {
		s.Logger.Error("Failed to get users",
			"pr_id", req.PullRequestId,
			"error", err)
		return nil, err
	}

	if model.Status == "MERGED" {
		s.Logger.Info("PR already merged")
		prDto := s.PrFromModelToDto(model)
		prDto.MergedAt = *model.MergedAt
		prDto.Reviewers = users
		return prDto, nil
	}

	now := time.Now().Truncate(time.Second)
	err = s.IPullRequest.SetMergeStatus(req.PullRequestId, now)
	if err != nil {
		s.Logger.Error("Failed to make merge",
			"error", err)
		return nil, err
	}

	newModel, err := s.IPullRequest.GetPullRequestById(req.PullRequestId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("PR not found")
			return nil, apperrors.ErrPrNotFound
		}
		s.Logger.Error("Failed to get PR",
			"error", err)
		return nil, err
	}

	dtoPr := s.PrFromModelToDto(newModel)
	dtoPr.MergedAt = *newModel.MergedAt
	dtoPr.Reviewers = users

	s.Logger.Info("Merge status set successfully",
		"pr_id", req.PullRequestId,
		"merged_at", now)

	return dtoPr, nil
}

func (s *Service) ReassignReviewer(req dto.ReassignReviewerReq) (*dto.PullRequestResp, error) {
	s.Logger.Info("Starting function ReassignReviewer",
		"pr_id", req.PullRequestId,
		"user_id", req.UserId)
	pr, err := s.IPullRequest.GetPullRequestById(req.PullRequestId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("PR not found",
				"pr_id", req.PullRequestId)
			return nil, apperrors.ErrPrNotFound
		}
		s.Logger.Error("Failed to get PR",
			"error", err)
		return nil, err
	}

	if pr.Status == "MERGED" {
		s.Logger.Warn("Cannot reassigne merged pr")
		return nil, apperrors.ErrPrMerged
	}

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
	s.Logger.Debug("Get user",
		"user_id", req.UserId)

	isAssigneg, err := s.IPullRequest.GetPullRequestByUser(req.UserId, req.PullRequestId)
	if err != nil {
		s.Logger.Error("Failed to get PR from user",
			"error", err)
		return nil, err
	}
	if !isAssigneg {
		s.Logger.Warn("PR not assigned to user",
			"pr_id", req.PullRequestId,
			"user_id", req.UserId)
		return nil, apperrors.ErrNotAssigned
	}
	s.Logger.Debug("Get PR by user")

	newAssigneeId, err := s.IUserRepo.GetUserToPr(user.TeamName, pr.AuthorId, pr.PullRequestId, req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("No candidate to asignee")
			return nil, apperrors.ErrNoCandidate
		}
		s.Logger.Error("Failed to find candidate",
			"error", err)
		return nil, err
	}
	s.Logger.Info("new assignee id", "id", newAssigneeId)

	err = s.IUserRepo.ChangeUserAssignee(req.PullRequestId, req.UserId, newAssigneeId)
	if err != nil {
		s.Logger.Error("Failed to change user assignee",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Change assignee successfully")

	prDto := s.PrFromModelToDto(pr)
	newUsers, err := s.IUserRepo.GetUsersByPr(req.PullRequestId)
	if err != nil {
		s.Logger.Error("Failed to get users",
			"error", err)
		return nil, err
	}
	prDto.Reviewers = newUsers
	s.Logger.Info("Change reviewer successfully",
		"old reviewer id", req.UserId,
		"new reviewer id", newAssigneeId)

	return prDto, nil
}

func (s *Service) GetPullRequestsByUser(userId string) ([]dto.PullRequestResp, error) {
	s.Logger.Info("Starting function: GetPullRequestsByUser",
		"user_id", userId)
	_, err := s.IUserRepo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Logger.Warn("User not found",
				"user_id", userId)
			return nil, apperrors.ErrUserNotFound
		}
		s.Logger.Error("Failed to get user",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Take user")

	pullRequests, err := s.IPullRequest.GetPullRequestsByUser(userId)
	if err != nil {
		s.Logger.Error("Failed to get PRs of user",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Take PRs")

	var dtoPullRequests []dto.PullRequestResp

	for _, pr := range pullRequests {
		dtoPr := s.PrFromModelToDto(&pr)
		dtoPullRequests = append(dtoPullRequests, *dtoPr)
	}

	s.Logger.Info("Successfully take all PRs of user",
		"user_id", userId,
		"count", len(pullRequests))

	return dtoPullRequests, nil
}

func (s *Service) PrFromModelToDto(req *models.PullRequestModel) *dto.PullRequestResp {
	return &dto.PullRequestResp{
		PullRequestId:   req.PullRequestId,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorId,
		Status:          req.Status,
	}
}

func (s *Service) PrFromDtoToModel(req *dto.CreatePullRequestReq) *models.PullRequestModel {
	return &models.PullRequestModel{
		PullRequestId:   req.PullRequestId,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorId,
	}
}
