package user

import (
	"database/sql"
	"log/slog"
)

type UserRepoSQL struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewUserRepoSQL(db *sql.DB, logger *slog.Logger) *UserRepoSQL {
	return &UserRepoSQL{
		db:     db,
		logger: logger,
	}
}
