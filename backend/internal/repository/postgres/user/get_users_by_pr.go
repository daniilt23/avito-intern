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
	defer rows.Close()

	var userId []string

	for rows.Next() {
		var user models.UserModel
		if err := rows.Scan(&user.UserId); err != nil {
			return nil, err
		}
		userId = append(userId, user.UserId)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userId, nil
}
