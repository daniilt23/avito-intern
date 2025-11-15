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
	defer rows.Close()

	var usersId []string

	for rows.Next() {
		var user models.UserModel
		if err := rows.Scan(&user.UserId); err != nil {
			return nil, err
		}
		usersId = append(usersId, user.UserId)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersId, nil
}
