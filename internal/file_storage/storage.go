package filestorage

import (
	"context"

	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

const ACCEPTED_CONTENT_TYPE = "image/png"

type FilesStorage interface {
	UploadFile(ctx context.Context, object files_portal.FileItem) (string, error)
}
