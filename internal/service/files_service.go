package service

import (
	"context"
	"time"

	filestorage "github.com/salesforceanton/files-portal/internal/file_storage"
	"github.com/salesforceanton/files-portal/internal/repository"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

type FilesService struct {
	repo    repository.Files
	storage filestorage.FilesStorage
}

func NewFilesService(repo repository.Files, storage filestorage.FilesStorage) *FilesService {
	return &FilesService{
		repo:    repo,
		storage: storage,
	}
}

func (s *FilesService) AddFileInfo(file files_portal.FileItem, userId int) (int, error) {
	// Put file to filestorage
	fileUrl, err := s.storage.UploadFile(context.TODO(), file)
	if err != nil {
		return 0, err
	}

	// Put file info into db
	return s.repo.AddFileInfo(files_portal.File{
		Size:        int(file.Size),
		Url:         fileUrl,
		Name:        file.Filename,
		CreatedDate: time.Now().UTC().String(),
		OwnerId:     userId,
	})
}
func (s *FilesService) GetFiles(userId int) ([]files_portal.File, error) {
	return s.repo.GetFiles(userId)
}
