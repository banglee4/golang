package repository

import (
	"sippm/model/domain"

	"gorm.io/gorm"
)

type PengDokumenRepository interface {
	Update(pengumumanId int, dokumen domain.DokumenPengumuman, db *gorm.DB) error
	DeleteByPengumumanId(id int, db *gorm.DB) error
}
