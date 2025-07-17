package service

import (
	"mime/multipart"
	"sippm/model/domain"
)

type FileService interface {
	UploadFile(files []*multipart.FileHeader) ([]domain.File, error)
	DeleteFile(fileURL string) error
}
