package team

import "avito/internal/models"

func (r *TeamRepoSQL) CreateTeam(reqTeam *models.TeamModel, reqUsers []models.UserModel) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			r.logger.Error("failed to rollback",
				"error", err)
		}
	}()

	query := `
	INSERT INTO teams (team_name) 
	VALUES ($1)`

	_, err = tx.Exec(query, reqTeam.TeamName)
	if err != nil {
		return err
	}

	query = `
	INSERT INTO users (user_id, username, is_active, team_name)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (user_id) DO UPDATE
	SET username = EXCLUDED.username, is_active = EXCLUDED.is_active, team_name = EXCLUDED.team_name`

	for _, user := range reqUsers {
		_, err := tx.Exec(query, user.UserId, user.Username, user.IsActive, user.TeamName)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
