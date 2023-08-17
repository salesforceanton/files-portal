package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user files_portal.User) (int, error) {
	var result int

	query := fmt.Sprintf(
		`INSERT INTO %s (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id`,
		USERS_TABLE,
	)
	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *AuthPostgres) GetUser(username, password string) (files_portal.User, error) {
	var result files_portal.User

	query := fmt.Sprintf(
		`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`,
		USERS_TABLE,
	)
	err := r.db.Get(&result, query, username, password)

	return result, err
}
