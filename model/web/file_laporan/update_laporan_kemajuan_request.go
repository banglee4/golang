package file_laporan

import "sippm/model/domain"

type UpdateFileLaporanKemajuanRequest struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func UpdateFileLaporanKemajuanRequestToDomain(request UpdateFileLaporanKemajuanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:                  request.Id,
		FileLaporanKemajuan: request.Url,
	}
}

type UpdateCatatanLaporanKemajuanRequest struct {
	Id      int    `json:"id"`
	Catatan string `json:"catatan"`
}

func UpdateCatatanLaporanKemajuanRequestToDomain(request UpdateCatatanLaporanKemajuanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:                     request.Id,
		CatatanLaporanKemajuan: request.Catatan,
	}
}

type UpdateStatusLaporanKemajuanRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func UpdateStatusLaporanKemajuanRequestToDomain(request UpdateStatusLaporanKemajuanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:                    request.Id,
		StatusLaporanKemajuan: request.Status,
	}
}
