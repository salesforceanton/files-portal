package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type FilesPostgres struct {
	db *sqlx.DB
}

func NewFilesPostgres(db *sqlx.DB) *FilesPostgres {
	return &FilesPostgres{db: db}
}

func (r *FilesPostgres) AddFileInfo(file files_portal.File) (int, error) {
	var result int

	query := fmt.Sprintf(
		`INSERT INTO %s (size, url, owner_id, created_date) VALUES ($1, $2, $3, $4) RETURNING id`,
		FILES_TABLE,
	)
	row := r.db.QueryRow(query, file.Size, file.Url, file.OwnerId, time.Now().UTC().String())

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *FilesPostgres) GetFiles(userId int) ([]files_portal.File, error) {
	var result []files_portal.File

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE owner_id=$1`,
		FILES_TABLE,
	)
	err := r.db.Select(&result, query, userId)

	return result, err
}
