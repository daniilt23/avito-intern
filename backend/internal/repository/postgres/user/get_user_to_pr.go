package user

func (r *UserRepoSQL) GetUserToPr(teamName string, authorId string, prId string, oldId string) (string, error) {
	query := `
	SELECT user_id FROM users u
	WHERE u.team_name = $1 AND u.is_active = TRUE AND u.user_id != $2 AND u.user_id != $3
	AND u.user_id NOT IN (
		SELECT pr.user_id FROM pr_users pr
		WHERE pr.pull_request_id = $4)
	ORDER BY RANDOM()
	LIMIT 1`

	var userId string
	err := r.db.QueryRow(query, teamName, authorId, oldId, prId).Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}