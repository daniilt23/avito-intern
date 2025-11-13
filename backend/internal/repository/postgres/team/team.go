package team

import "database/sql"

type TeamRepoSQL struct {
	db *sql.DB
}

func NewTeamRepoSQL(db *sql.DB) *TeamRepoSQL {
	return &TeamRepoSQL{db: db}
}