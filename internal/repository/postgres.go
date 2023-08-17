package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/salesforceanton/files-portal/internal/config"
)

const (
	POSTGRESS_DB_TYPE = "postgres"
	USERS_TABLE       = "users"
	FILES_TABLE       = "files"
)

func NewPostgresDB(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", POSTGRESS_DB_TYPE, cfg.Username, cfg.Password, cfg.Host, cfg.Name))
	db, err := sqlx.Open(POSTGRESS_DB_TYPE, pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
