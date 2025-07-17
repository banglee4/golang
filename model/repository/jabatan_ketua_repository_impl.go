package repository

import (
	"fmt"
	"gorm.io/gorm"
	"sippm/model/domain"
)

type JabatanRepositoryImpl struct {
}

func NewJabatanRepositoryImpl() *JabatanRepositoryImpl {
	return &JabatanRepositoryImpl{}
}

func (repo JabatanRepositoryImpl) Create(ketua domain.Ketua, db *gorm.DB) error {
	err := db.Create(&ketua).Error
	if err != nil {
		return fmt.Errorf("failed to create ketua: %w", err)
	}
	return nil
}

func (repo JabatanRepositoryImpl) Update(ketua domain.Ketua, db *gorm.DB) error {
	err := db.Model(&domain.Ketua{}).
		Where("id = ?", ketua.Id).
		Updates(map[string]interface{}{
			"nama_dosen": ketua.NamaDosen,
			"jabatan":    ketua.Jabatan,
		}).Error

	if err != nil {
		return fmt.Errorf("failed to update jabatan: %w", err)
	}

	return nil
}

func (repo JabatanRepositoryImpl) GetAll(db *gorm.DB) ([]domain.Ketua, error) {
	var results []domain.Ketua

	err := db.
		Select("id", "nama_dosen", "jabatan").
		Order("created_at DESC").
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get ketua: %w", err)
	}

	return results, nil
}

func (repo JabatanRepositoryImpl) Delete(id int, db *gorm.DB) error {
	err := db.Delete(&domain.Ketua{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete ketua: %w", err)
	}
	return nil
}
