package filestorage

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/salesforceanton/files-portal/internal/config"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

const PHOTO_BUCKET_NAME = "photos"

type MinioProvider struct {
	client     *minio.Client
	url        string
	bucketName string
}

func NewMinioProvider(cfg config.MinioConfig, bucketName string) (*MinioProvider, error) {
	// Create minio client
	minioUrl := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	fmt.Println(minioUrl)
	client, err := minio.New(minioUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Username, cfg.Password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	// Create bucket
	err = client.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return nil, err
	}

	return &MinioProvider{
		client: client,
		url:    minioUrl,
	}, nil
}

func (m *MinioProvider) UploadFile(ctx context.Context, object files_portal.FileItem) (string, error) {
	filename := m.GenerateFileName(object.Filename)

	_, err := m.client.PutObject(
		ctx,
		m.bucketName,
		filename,
		object.Source,
		object.Size,
		minio.PutObjectOptions{ContentType: ACCEPTED_CONTENT_TYPE},
	)

	return m.GenerateFileUrl(filename), err
}

func (m *MinioProvider) GenerateFileUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s", m.url, m.bucketName, filename)
}

func (m *MinioProvider) GenerateFileName(filename string) string {
	return fmt.Sprintf("%s-%s", filename, time.Now().UTC().String())
}
