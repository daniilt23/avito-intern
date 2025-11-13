package user

import "database/sql"

type UserRepoSQL struct {
	db *sql.DB
}

func NewUserRepoSQL(db *sql.DB) *UserRepoSQL {
	return &UserRepoSQL{db: db}
}

func (r *UserRepoSQL) CreateUser(name string, isActive bool, teamId int) error {
	return nil
}