package team

func (r *TeamRepoSQL) GetTeamByName(name string) (string, error) {
	query := `
	SELECT team_name FROM teams
	WHERE team_name = $1`

	var team_name string
	if err := r.db.QueryRow(query, name).Scan(&team_name); err != nil {
		return "", err
	}

	return team_name, nil
}
