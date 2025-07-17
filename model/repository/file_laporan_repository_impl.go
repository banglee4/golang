package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type FileLaporanRepositoryImpl struct {
	DB *gorm.DB
}

func NewFileLaporanRepositoryImpl(DB *gorm.DB) *FileLaporanRepositoryImpl {
	return &FileLaporanRepositoryImpl{DB: DB}
}

func (repo *FileLaporanRepositoryImpl) UpdateFileUsulan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"file_usulan": file.FileUsulan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update file usulan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateCatatanUsulan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"catatan_usulan": file.CatatanUsulan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update catatan usulan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateFileLaporanKemajuan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"file_laporan_kemajuan": file.FileLaporanKemajuan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update file usulan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateCatatanLaporanKemajuan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"catatan_laporan_kemajuan": file.CatatanLaporanKemajuan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update catatan laporan kemajuan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateFileLaporanAkhir(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"file_laporan_akhir": file.FileLaporanAkhir,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update file laporan akhir")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateCatatanLaporanAkhir(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"catatan_laporan_akhir": file.CatatanLaporanAkhir,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update catatan laporan akhir")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateStatusUsulan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"status_usulan": file.StatusUsulan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update status usulan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateStatusLaporanKemajuan(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"status_laporan_kemajuan": file.StatusLaporanKemajuan,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update status laporan kemajuan")
	}

	return nil
}

func (repo *FileLaporanRepositoryImpl) UpdateStatusLaporanAkhir(file domain.FileLaporan) error {
	err := repo.DB.
		Model(&file).
		Where("id = ?", file.Id).
		Updates(map[string]interface{}{
			"status_laporan_akhir": file.StatusLaporanAkhir,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update status laporan kemajuan")
	}

	return nil
}
