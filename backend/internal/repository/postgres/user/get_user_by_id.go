package user

import "avito/internal/models"

func (r *UserRepoSQL) GetUserById(id string) (models.UserModel, error) {
	query := `SELECT user_id, username, is_active, team_name
	FROM users WHERE user_id = $1`

	model := models.UserModel{}
	err := r.db.QueryRow(query, id).Scan(&model.UserId, &model.Username, &model.IsActive, &model.TeamName)
	if err != nil {
		return model, err
	}

	return model, nil
}