package repository

import (
	"gorm.io/gorm"
	"sippm/model/domain"
)

type JabatanKetuaRepository interface {
	Create(ketua domain.Ketua, db *gorm.DB) error
	Update(ketua domain.Ketua, db *gorm.DB) error
	GetAll(db *gorm.DB) ([]domain.Ketua, error)
	Delete(id int, db *gorm.DB) error
}
