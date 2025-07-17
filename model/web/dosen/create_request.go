package dosen

import "sippm/model/domain"

type CreateRequest struct {
	Nidn  string `json:"nidn"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func CreateRequestToDomain(request CreateRequest) domain.Dosen {
	return domain.Dosen{
		Nidn:  request.Nidn,
		Nama:  request.Nama,
		Email: request.Email,
		Phone: request.Phone,
	}
}
