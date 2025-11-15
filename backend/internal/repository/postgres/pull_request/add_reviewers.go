package pullrequest

func (r *PullRequestRepoSQL) AddReviewers(usersId []string, prId string) error {
	if len(usersId) == 0 {
		return nil
	}

	query := `INSERT INTO pr_users (pull_request_id, user_id) 
	VALUES ($1, $2)`

	for _, userId := range usersId {
		_, err := r.db.Exec(query, prId, userId)
		if err != nil {
			return err
		}
	}

	return nil
}
