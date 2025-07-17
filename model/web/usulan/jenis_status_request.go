package usulan

import (
	"sippm/model/domain"
)

type JenisStatusRequest struct {
	Jenis  string `json:"jenis"`
	Status string `json:"status"`
}

func JenisStatusRequestToDomain(request JenisStatusRequest) domain.Usulan {
	return domain.Usulan{
		Jenis:  request.Jenis,
		Status: request.Status,
	}
}
