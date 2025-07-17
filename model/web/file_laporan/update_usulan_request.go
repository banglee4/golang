package file_laporan

import "sippm/model/domain"

type UpdateFileUsulanRequest struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func UpdateFileUsulanRequestToDomain(request UpdateFileUsulanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:         request.Id,
		FileUsulan: request.Url,
	}
}

type UpdateCatatanUsulanRequest struct {
	Id      int    `json:"id"`
	Catatan string `json:"catatan"`
}

func UpdateCatatanUsulanRequestToDomain(request UpdateCatatanUsulanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:            request.Id,
		CatatanUsulan: request.Catatan,
	}
}

type UpdateStatusUsulanRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func UpdateStatusUsulanRequestToDomain(request UpdateStatusUsulanRequest) domain.FileLaporan {
	return domain.FileLaporan{
		Id:           request.Id,
		StatusUsulan: request.Status,
	}
}
