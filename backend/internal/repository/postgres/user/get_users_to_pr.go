package user

import "avito/internal/models"

func (r *UserRepoSQL) GetUsersToPr(teamName string, authorId string) ([]string, error) {
	query := `
	SELECT user_id FROM users
	WHERE team_name = $1 AND is_active = TRUE AND user_id != $2
	ORDER BY RANDOM()
	LIMIT 2`

	rows, err := r.db.Query(query, teamName, authorId)
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

	var usersId []string

	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.UserId)
		if err != nil {
			return nil, err
		}
		usersId = append(usersId, user.UserId)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usersId, nil
}
