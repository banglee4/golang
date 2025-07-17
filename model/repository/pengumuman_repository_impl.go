package repository

import (
	"fmt"
	"gorm.io/gorm"
	"sippm/model/domain"
)

type PengumumanRepositoryImpl struct {
}

func NewPengumumanRepositoryImpl() *PengumumanRepositoryImpl {
	return &PengumumanRepositoryImpl{}
}

func (repo *PengumumanRepositoryImpl) Create(model domain.Pengumuman, db *gorm.DB) error {
	err := db.Create(&model).Error
	if err != nil {
		return fmt.Errorf("failed to create pengumuman: %w", err)
	}
	return nil
}

func (repo *PengumumanRepositoryImpl) Update(model domain.Pengumuman, db *gorm.DB) error {
	err := db.Model(&domain.Pengumuman{}).
		Where("id = ?", model.Id).
		Updates(map[string]interface{}{
			"judul":      model.Judul,
			"created_at": model.TanggalDibuat,
		}).Error

	if err != nil {
		return fmt.Errorf("failed to update pengumuman: %w", err)
	}

	return nil
}

func (repo *PengumumanRepositoryImpl) GetAll(db *gorm.DB) ([]domain.Pengumuman, error) {
	var results []domain.Pengumuman

	err := db.
		Select("id", "judul", "created_at").
		Preload("Dokumen").
		Order("created_at DESC").
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get dokumen: %w", err)
	}

	return results, nil
}

func (repo *PengumumanRepositoryImpl) Delete(id int, db *gorm.DB) error {
	err := db.Delete(&domain.Pengumuman{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete pengumuman: %w", err)
	}
	return nil
}
