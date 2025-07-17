package service

import (
	"sippm/model/web/file_laporan"
)

type FileLaporanService interface {
	UpdateFileUsulan(request file_laporan.UpdateFileUsulanRequest) error
	UpdateStatusUsulan(request file_laporan.UpdateStatusUsulanRequest) error
	UpdateCatatanUsulan(request file_laporan.UpdateCatatanUsulanRequest) error
	UpdateFileLaporanKemajuan(request file_laporan.UpdateFileLaporanKemajuanRequest) error
	UpdateStatusLaporanKemajuan(request file_laporan.UpdateStatusLaporanKemajuanRequest) error
	UpdateCatatanLaporanKemajuan(request file_laporan.UpdateCatatanLaporanKemajuanRequest) error
	UpdateFileLaporanAkhir(request file_laporan.UpdateFileLaporanAkhirRequest) error
	UpdateStatusLaporanAkhir(request file_laporan.UpdateStatusLaporanAkhirRequest) error
	UpdateCatatanLaporanAkhir(request file_laporan.UpdateCatatanLaporanAkhirRequest) error
}
