package pullrequest

import (
	"database/sql"
	"log/slog"
)

type PullRequestRepoSQL struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewPullRequestRepoSQL(db *sql.DB, logger *slog.Logger) *PullRequestRepoSQL {
	return &PullRequestRepoSQL{
		db: db,
		logger: logger,
	}
}
