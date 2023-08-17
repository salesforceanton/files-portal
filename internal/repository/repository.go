package repository

import (
	"github.com/jmoiron/sqlx"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type Repository struct {
	Authorization
	Files
}

type Authorization interface {
	CreateUser(user files_portal.User) (int, error)
	GetUser(username, password string) (files_portal.User, error)
}

type Files interface {
	AddFileInfo(file files_portal.File) (int, error)
	GetFiles(userId int) ([]files_portal.File, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Files:         NewFilesPostgres(db),
	}
}
