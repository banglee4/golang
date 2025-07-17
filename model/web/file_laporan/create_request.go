package file_laporan

import "sippm/model/domain"

type CreateRequest struct {
	FileUsulan string `json:"file_usulan"`
}

func CreateRequestToDomain(request CreateRequest) domain.FileLaporan {
	return domain.FileLaporan{
		FileUsulan:   request.FileUsulan,
		StatusUsulan: "Pending",
	}
}
