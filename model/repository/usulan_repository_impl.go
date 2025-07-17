package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type UsulanRepositoryImpl struct {
	DB *gorm.DB
}

func NewUsulanRepositoryImpl(DB *gorm.DB) *UsulanRepositoryImpl {
	return &UsulanRepositoryImpl{DB: DB}
}

func (repo *UsulanRepositoryImpl) Create(usulan domain.Usulan) error {
	err := repo.DB.Create(&usulan).Error
	if err != nil {
		return fmt.Errorf("failed to create usulan: %w", err)
	}
	return nil
}

func (repo *UsulanRepositoryImpl) Update(usulan domain.Usulan) error {
	err := repo.DB.
		Model(&usulan).
		Where("id = ?", usulan.Id).
		Updates(map[string]interface{}{
			"judul":            usulan.Judul,
			"fakultas":         usulan.Fakultas,
			"prodi":            usulan.Prodi,
			"jenis_usulan":     usulan.Jenis,
			"tahun_usulan":     usulan.Tahun,
			"anggaran":         usulan.Anggaran,
			"nama_anggota_1":   usulan.NamaAnggota1,
			"nama_anggota_2":   usulan.NamaAnggota2,
			"nama_anggota_3":   usulan.NamaAnggota3,
			"nama_mahasiswa_1": usulan.NamaMahasiswa1,
			"nama_mahasiswa_2": usulan.NamaMahasiswa2,
			"nama_mahasiswa_3": usulan.NamaMahasiswa3,
			"nama_mahasiswa_4": usulan.NamaMahasiswa4,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update usulan")
	}

	return nil
}

func (repo *UsulanRepositoryImpl) GetByJenisAndStatus(usulan domain.Usulan) ([]domain.Usulan, int, error) {
	var results []domain.Usulan
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Usulan{}).Where("jenis_usulan = ? AND status = ?", usulan.Jenis, usulan.Status).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count usulan: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "anggaran", "status", "nama_anggota_1", "nama_anggota_2", "nama_anggota_3", "nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen").
		Where("jenis_usulan = ? AND status = ?", usulan.Jenis, usulan.Status).
		Order("updated_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all usulan: %w", err)
	}

	return results, int(total), nil
}

func (repo *UsulanRepositoryImpl) SearchByJudul(usulan domain.Usulan) ([]domain.Usulan, int, error) {
	var results []domain.Usulan
	var total int64

	if err := repo.DB.Model(&domain.Usulan{}).
		Where("judul LIKE ?", "%"+usulan.Judul+"%").
		Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count usulan: %w", err)
	}

	err := repo.DB.
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "status", "nama_anggota_1", "nama_anggota_2", "nama_anggota_3", "nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen").
		Where("judul LIKE ?", "%"+usulan.Judul+"%").
		Order("updated_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to search usulan by judul: %w", err)
	}

	return results, int(total), nil
}

func (repo *UsulanRepositoryImpl) Filter(usulan domain.Usulan, tahunMulai, tahunAkhir string) ([]domain.Usulan, error) {
	var results []domain.Usulan

	query := repo.DB.Model(&domain.Usulan{}).
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "status",
			"nama_anggota_1", "nama_anggota_2", "nama_anggota_3",
			"nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen")

	// Filter Prodi
	if usulan.Prodi != "" {
		query = query.Where("prodi = ?", usulan.Prodi)
	}

	// Filter Jenis Usulan
	if usulan.Jenis != "" {
		query = query.Where("jenis_usulan = ?", usulan.Jenis)
	}

	if tahunMulai != "" && tahunAkhir != "" {
		query = query.Where("tahun_usulan BETWEEN ? AND ?", tahunMulai, tahunAkhir)
	} else if tahunMulai != "" {
		query = query.Where("tahun_usulan >= ?", tahunMulai)
	} else if tahunAkhir != "" {
		query = query.Where("tahun_usulan <= ?", tahunAkhir)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to filter usulan: %w", err)
	}

	return results, nil
}

func (repo *UsulanRepositoryImpl) UpdateStatus(usulan domain.Usulan) error {
	err := repo.DB.
		Model(&usulan).
		Where("id = ?", usulan.Id).
		Updates(map[string]interface{}{
			"status": usulan.Status,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update status usulan")
	}

	return nil
}

func (repo *UsulanRepositoryImpl) GetAll(jenis string, fakultas string) ([]domain.Usulan, int, error) {
	var results []domain.Usulan
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Usulan{}).Where("jenis_usulan = ? AND fakultas = ?", jenis, fakultas).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count usulan: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "anggaran", "status", "nama_anggota_1", "nama_anggota_2", "nama_anggota_3", "nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen").
		Preload("FileLaporan").
		Where("status != 'Draft' AND jenis_usulan = ? AND fakultas = ?", jenis, fakultas).
		Order("updated_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all usulan: %w", err)
	}

	return results, int(total), nil
}

func (repo *UsulanRepositoryImpl) GetAllValidated(jenis string) ([]domain.Usulan, int, error) {
	var results []domain.Usulan
	var total int64

	commonConditions := "status != 'Draft' AND status != 'Usulan Pending' AND status != 'Usulan Ditolak' AND jenis_usulan = ?"

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Usulan{}).Where(commonConditions, jenis).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count usulan: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "anggaran", "status", "nama_anggota_1", "nama_anggota_2", "nama_anggota_3", "nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen").
		Preload("FileLaporan").
		Where(commonConditions, jenis).
		Order("updated_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all usulan: %w", err)
	}

	return results, int(total), nil
}

func (repo *UsulanRepositoryImpl) GetAllByPengusulId(pengusulId int, jenis string) ([]domain.Usulan, int, error) {
	var results []domain.Usulan
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Usulan{}).Where("pengusul_id = ?", pengusulId).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count usulan: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "pengusul_id", "fakultas", "prodi", "judul", "jenis_usulan", "tahun_usulan", "anggaran", "status", "nama_anggota_1", "nama_anggota_2", "nama_anggota_3", "nama_mahasiswa_1", "nama_mahasiswa_2", "nama_mahasiswa_3", "nama_mahasiswa_4").
		Preload("User").
		Preload("User.Dosen").
		Preload("FileLaporan").
		Where("pengusul_id = ? AND jenis_usulan = ?", pengusulId, jenis).
		Order("updated_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all usulan: %w", err)
	}

	return results, int(total), nil
}

func (repo *UsulanRepositoryImpl) Delete(id int) error {
	// Hapus data dari FileLaporan yang terkait dengan usulan
	if err := repo.DB.Where("usulan_id = ?", id).Delete(&domain.FileLaporan{}).Error; err != nil {
		return fmt.Errorf("failed to delete related file laporan: %w", err)
	}

	// Hapus data dari Usulan
	if err := repo.DB.Delete(&domain.Usulan{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete usulan: %w", err)
	}

	return nil
}
