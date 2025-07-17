package repository

import (
	"sippm/model/domain"

	"gorm.io/gorm"
)

type PengumumanRepository interface {
	Create(model domain.Pengumuman, db *gorm.DB) error
	Update(model domain.Pengumuman, db *gorm.DB) error
	GetAll(db *gorm.DB) ([]domain.Pengumuman, error)
	Delete(id int, db *gorm.DB) error
}
