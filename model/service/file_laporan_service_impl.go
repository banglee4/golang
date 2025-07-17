package service

import (
	"sippm/model/repository"
	"sippm/model/web/file_laporan"
)

type FileLaporanServiceImpl struct {
	Repo repository.FileLaporanRepository
}

func NewFileLaporanServiceImpl(repo repository.FileLaporanRepository) *FileLaporanServiceImpl {
	return &FileLaporanServiceImpl{Repo: repo}
}

func (serv *FileLaporanServiceImpl) UpdateFileUsulan(request file_laporan.UpdateFileUsulanRequest) error {
	model := file_laporan.UpdateFileUsulanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateFileUsulan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateCatatanUsulan(request file_laporan.UpdateCatatanUsulanRequest) error {
	model := file_laporan.UpdateCatatanUsulanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateCatatanUsulan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateFileLaporanKemajuan(request file_laporan.UpdateFileLaporanKemajuanRequest) error {
	model := file_laporan.UpdateFileLaporanKemajuanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateFileLaporanKemajuan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateCatatanLaporanKemajuan(request file_laporan.UpdateCatatanLaporanKemajuanRequest) error {
	model := file_laporan.UpdateCatatanLaporanKemajuanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateCatatanLaporanKemajuan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateFileLaporanAkhir(request file_laporan.UpdateFileLaporanAkhirRequest) error {
	model := file_laporan.UpdateFileLaporanAkhirRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateFileLaporanAkhir(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateCatatanLaporanAkhir(request file_laporan.UpdateCatatanLaporanAkhirRequest) error {
	model := file_laporan.UpdateCatatanLaporanAkhirRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateCatatanLaporanAkhir(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateStatusUsulan(request file_laporan.UpdateStatusUsulanRequest) error {
	model := file_laporan.UpdateStatusUsulanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateStatusUsulan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateStatusLaporanKemajuan(request file_laporan.UpdateStatusLaporanKemajuanRequest) error {
	model := file_laporan.UpdateStatusLaporanKemajuanRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateStatusLaporanKemajuan(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *FileLaporanServiceImpl) UpdateStatusLaporanAkhir(request file_laporan.UpdateStatusLaporanAkhirRequest) error {
	model := file_laporan.UpdateStatusLaporanAkhirRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateStatusLaporanAkhir(model)
	if err != nil {
		return err
	}

	return nil
}
