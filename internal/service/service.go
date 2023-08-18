package service

import (
	"github.com/salesforceanton/files-portal/internal/config"
	"github.com/salesforceanton/files-portal/internal/repository"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type Service struct {
	Authorization
	Files
}

type Authorization interface {
	CreateUser(user files_portal.User) (int, error)
	GenerateAccesssToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Files interface {
	AddFileInfo(file files_portal.File) (int, error)
	GetFiles(userId int) ([]files_portal.File, error)
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Files:         NewFilesService(repos.Files),
	}
}
