package pullrequest

import (
	"avito/internal/models"
	"database/sql"
)

func (r *PullRequestRepoSQL) GetPullRequestById(id string) (*models.PullRequestModel, error) {
	query := `
	SELECT pull_request_id, pull_request_name, author_id, status, merged_at FROM pull_requests
	WHERE pull_request_id = $1`

	var mergedAt sql.NullTime

	var model models.PullRequestModel
	err := r.db.QueryRow(query, id).Scan(&model.PullRequestId, &model.PullRequestName,
		&model.AuthorId, &model.Status, &mergedAt)
	if err != nil {
		return nil, err
	}

	if mergedAt.Valid {
		model.MergedAt = &mergedAt.Time
	}

	return &model, nil
}
