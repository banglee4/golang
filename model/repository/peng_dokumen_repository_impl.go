package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type PengDokumenRepositoryImpl struct {
}

func NewPengDokumenRepositoryImpl() *PengDokumenRepositoryImpl {
	return &PengDokumenRepositoryImpl{}
}

func (repo *PengDokumenRepositoryImpl) Update(pengumumanId int, dokumen domain.DokumenPengumuman, db *gorm.DB) error {
	if err := db.Create(&domain.DokumenPengumuman{
		NamaDokumen:  dokumen.NamaDokumen,
		Url:          dokumen.Url,
		PengumumanId: pengumumanId,
	}).Error; err != nil {
		return fmt.Errorf("gagal menyimpan dokumen baru: %w", err)
	}

	return nil
}

func (repo *PengDokumenRepositoryImpl) DeleteByPengumumanId(id int, db *gorm.DB) error {
	err := db.Where("pengumuman_id = ?", id).Delete(&domain.DokumenPengumuman{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete pengumuman dokumen by PengumumanId: %w", err)
	}
	return nil
}
