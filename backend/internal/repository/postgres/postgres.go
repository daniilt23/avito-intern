package postgres

import (
	"avito/internal/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

func ConnectDB(config config.Storage) (*sql.DB, error) {
	sdn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.DbName, config.Username, config.Password, config.Sslmode)

	db, err := sql.Open("postgres", sdn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
		return nil, err
	}

	err = goose.Up(db, "/app/migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}
