package pullrequest

import "avito/internal/models"

func (r *PullRequestRepoSQL) CreatePullRequest(model *models.PullRequestModel) error {
	query := `INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id)
	VALUES ($1, $2, $3)`

	_, err := r.db.Exec(query, model.PullRequestId, model.PullRequestName, model.AuthorId)
	if err != nil {
		return err
	}

	return nil
}