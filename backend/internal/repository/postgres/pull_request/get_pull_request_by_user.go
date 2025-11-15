package pullrequest

func (r *PullRequestRepoSQL) GetPullRequestByUser(userId string, pullRequestId string) (bool, error) {
	query := `
	SELECT COUNT(*) FROM pr_users
	WHERE user_id = $1 AND pull_request_id = $2
	LIMIT 1`

	var count int

	err := r.db.QueryRow(query, userId, pullRequestId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}