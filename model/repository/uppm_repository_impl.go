package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type UppmRepositoryImpl struct {
	DB *gorm.DB
}

func NewUppmRepositoryImpl(DB *gorm.DB) *UppmRepositoryImpl {
	return &UppmRepositoryImpl{DB: DB}
}

func (repo UppmRepositoryImpl) UpdateData(uppm domain.Uppm) error {
	err := repo.DB.
		Model(&uppm).
		Where("id = ?", uppm.Id).
		Updates(map[string]interface{}{
			"nama":     uppm.Nama,
			"photo":    uppm.Photo,
			"email":    uppm.Email,
			"fakultas": uppm.Fakultas,
			"phone":    uppm.Phone,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update data")
	}

	return nil
}
