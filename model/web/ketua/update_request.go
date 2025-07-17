package ketua

import "sippm/model/domain"

type UpdateRequest struct {
	Id        int    `json:"id"`
	NamaDosen string `json:"nama_dosen"`
	Jabatan   string `json:"jabatan"`
}

func UpdateReqToDomain(request UpdateRequest) domain.Ketua {
	return domain.Ketua{
		Id:        request.Id,
		NamaDosen: request.NamaDosen,
		Jabatan:   request.Jabatan,
	}
}
