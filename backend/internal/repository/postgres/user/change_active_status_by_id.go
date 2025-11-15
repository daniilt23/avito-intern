package user


func (r *UserRepoSQL) ChangeActiveStatusById(id string, isActive bool) error {
	query := `
	UPDATE users
	SET is_active = $1
	WHERE user_id = $2`

	_, err := r.db.Exec(query, isActive, id)
	if err != nil {
		return err
	}

	return nil
}