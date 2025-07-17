package file_laporan

import "sippm/model/domain"

type UpdateFileLaporanAkhirRequest struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func UpdateFileLaporanAkhirRequestToDomain(request UpdateFileLaporanAkhirRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:               request.Id,
		FileLaporanAkhir: request.Url,
	}
}

type UpdateCatatanLaporanAkhirRequest struct {
	Id      int    `json:"id"`
	Catatan string `json:"catatan"`
}

func UpdateCatatanLaporanAkhirRequestToDomain(request UpdateCatatanLaporanAkhirRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:                  request.Id,
		CatatanLaporanAkhir: request.Catatan,
	}
}

type UpdateStatusLaporanAkhirRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func UpdateStatusLaporanAkhirRequestToDomain(request UpdateStatusLaporanAkhirRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:                 request.Id,
		StatusLaporanAkhir: request.Status,
	}
}
