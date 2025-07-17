package uppm

import "sippm/model/domain"

type CreateRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fakultas string `json:"fakultas"`
}

func CreateRequestToDomain(request CreateRequest) domain.Uppm {
	return domain.Uppm{
		Nama:     request.Nama,
		Email:    request.Email,
		Phone:    request.Phone,
		Fakultas: request.Fakultas,
	}
}
