package pullrequest

import "avito/internal/models"

func (r *PullRequestRepoSQL) GetPullRequestsByUser(userId string) ([]models.PullRequestModel, error) {
	query := `
	SELECT pr.pull_request_id, pr.pull_request_name, pr.author_id, pr.status FROM pull_requests pr
	WHERE pr.pull_request_id IN (
		SELECT pru.pull_request_id FROM pr_users pru
		WHERE pru.user_id = $1)`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pullRequests []models.PullRequestModel

	for rows.Next() {
		var pullRequest models.PullRequestModel
		err := rows.Scan(&pullRequest.PullRequestId, &pullRequest.PullRequestName,
			&pullRequest.AuthorId, &pullRequest.Status)
		if err != nil {
			return nil, err
		}
		pullRequests = append(pullRequests, pullRequest)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return pullRequests, nil
}
