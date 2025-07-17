package uppm

import "sippm/model/domain"

type UpdateDataRequest struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Photo    string `json:"photo"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fakultas string `json:"fakultas"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domain.Uppm {
	return domain.Uppm{
		Id:       request.Id,
		Nama:     request.Nama,
		Photo:    request.Photo,
		Email:    request.Email,
		Phone:    request.Phone,
		Fakultas: request.Fakultas,
	}
}
