package repository

import "sippm/model/domain"

type FileLaporanRepository interface {
	UpdateFileUsulan(file domain.FileLaporan) error
	UpdateStatusUsulan(file domain.FileLaporan) error
	UpdateCatatanUsulan(file domain.FileLaporan) error
	UpdateStatusLaporanKemajuan(file domain.FileLaporan) error
	UpdateFileLaporanKemajuan(file domain.FileLaporan) error
	UpdateCatatanLaporanKemajuan(file domain.FileLaporan) error
	UpdateStatusLaporanAkhir(file domain.FileLaporan) error
	UpdateFileLaporanAkhir(file domain.FileLaporan) error
	UpdateCatatanLaporanAkhir(file domain.FileLaporan) error
}
