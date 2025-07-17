package ketua

import "sippm/model/domain"

type CreateRequest struct {
	NamaDosen string `json:"nama_dosen"`
	Jabatan   string `json:"jabatan"`
}

func CreateReqToDomain(request CreateRequest) domain.Ketua {
	return domain.Ketua{
		NamaDosen: request.NamaDosen,
		Jabatan:   request.Jabatan,
	}
}
