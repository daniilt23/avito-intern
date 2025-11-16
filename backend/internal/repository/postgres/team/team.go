package team

import (
	"database/sql"
	"log/slog"
)

type TeamRepoSQL struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewTeamRepoSQL(db *sql.DB, logger *slog.Logger) *TeamRepoSQL {
	return &TeamRepoSQL{
		db: db,
		logger: logger,	
	}
}
