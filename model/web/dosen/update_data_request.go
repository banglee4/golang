package dosen

import "sippm/model/domain"

type UpdateDataRequest struct {
	Id    int    `json:"id"`
	Nidn  string `json:"nidn"`
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domain.Dosen {
	return domain.Dosen{
		Id:    request.Id,
		Nidn:  request.Nidn,
		Nama:  request.Nama,
		Photo: request.Photo,
		Email: request.Email,
		Phone: request.Phone,
	}
}
