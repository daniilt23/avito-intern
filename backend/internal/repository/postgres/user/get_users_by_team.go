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
	defer func() {
		err := rows.Close()
		if err != nil {
			r.logger.Error("Failed to close rows",
				"error", err)
		}
	}()

	var users []models.UserModel

	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.UserId, &user.Username, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return &users, err
	}

	return &users, nil
}
