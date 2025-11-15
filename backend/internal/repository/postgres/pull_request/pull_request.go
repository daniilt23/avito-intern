package pullrequest

import "database/sql"

type PullRequestRepoSQL struct {
	db *sql.DB
}

func NewPullRequestRepoSQL(db *sql.DB) *PullRequestRepoSQL {
	return &PullRequestRepoSQL{
		db: db,
	}
}