package service

import files_portal "github.com/salesforceanton/files-portal/pkg/domain"

type Service struct {
	Authorization
	Files
}

type Authorization interface {
	CreateUser(user files_portal.User) (int, error)
	GenerateToken(user files_portal.User) (string, error)
	ParseToken(token string) error
}

type Files interface {
	AddFileInfo(file files_portal.File) (int, error)
	GetFiles(userId int) ([]files_portal.File, error)
}

// func NewService(repos repository.Repository) *Service {
// 	return &Service{
// 		Authorization: NewAuthPostgres(db),
// 		Files:         NewFilesPostgres(db),
// 	}
// }
