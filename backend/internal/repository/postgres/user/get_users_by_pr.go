package user

import "avito/internal/models"

func (r *UserRepoSQL) GetUsersByPr(prName string) ([]string, error) {
	query := `
	SELECT user_id FROM pr_users
	WHERE pull_request_id = $1`

	rows, err := r.db.Query(query, prName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			r.logger.Error("Failed to close rows",
				"error", err)
		}
	}()

	var userId []string

	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.UserId)
		if err != nil {
			return nil, err
		}
		userId = append(userId, user.UserId)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return userId, nil
}
