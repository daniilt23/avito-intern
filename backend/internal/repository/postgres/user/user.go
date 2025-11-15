package user

import (
	"database/sql"
)

type UserRepoSQL struct {
	db *sql.DB
}

func NewUserRepoSQL(db *sql.DB) *UserRepoSQL {
	return &UserRepoSQL{db: db}
}
