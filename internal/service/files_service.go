package service

import (
	"github.com/salesforceanton/files-portal/internal/repository"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type FilesService struct {
	repo repository.Files
}

func NewFilesService(repo repository.Files) *FilesService {
	return &FilesService{
		repo: repo,
	}
}

func (s *FilesService) AddFileInfo(file files_portal.File) (int, error) {
	return s.repo.AddFileInfo(file)
}
func (s *FilesService) GetFiles(userId int) ([]files_portal.File, error) {
	return s.repo.GetFiles(userId)
}
