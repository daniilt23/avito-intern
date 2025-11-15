package user

func (r *UserRepoSQL) ChangeUserAssignee(prId string, userId string, newUserId string) error {
	query := `
	UPDATE pr_users
	SET user_id = $1
	WHERE pull_request_id = $2 AND user_id = $3`

	_, err := r.db.Exec(query, newUserId, prId, userId)
	if err != nil {
		return err
	}

	return nil
}