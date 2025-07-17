package lppm

import "sippm/model/domain"

type CreateRequest struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func CreateRequestToDomain(request CreateRequest) domain.Lppm {
	return domain.Lppm{
		Nama:  request.Nama,
		Email: request.Email,
		Phone: request.Phone,
	}
}
