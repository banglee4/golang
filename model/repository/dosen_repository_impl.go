package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type DosenRepositoryImpl struct {
	DB *gorm.DB
}

func NewDosenRepositoryImpl(DB *gorm.DB) *DosenRepositoryImpl {
	return &DosenRepositoryImpl{DB: DB}
}

func (repo DosenRepositoryImpl) UpdateData(dosen domain.Dosen) error {
	err := repo.DB.
		Model(&dosen).
		Where("id = ?", dosen.Id).
		Updates(map[string]interface{}{
			"nidn":  dosen.Nidn,
			"nama":  dosen.Nama,
			"photo": dosen.Photo,
			"email": dosen.Email,
			"phone": dosen.Phone,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update data")
	}

	return nil
}
