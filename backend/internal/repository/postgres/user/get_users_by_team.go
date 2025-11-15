package user

import "avito/internal/models"

func (r *UserRepoSQL) GetUsersByTeam(name string) (*[]models.UserModel, error) {
	query := `
	SELECT user_id, username, is_active FROM users 
	WHERE team_name = $1`

	rows, err := r.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserModel

	for rows.Next() {
		var user models.UserModel
		if err := rows.Scan(&user.UserId, &user.Username, &user.IsActive); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return &users, err
	}

	return &users, nil
}
