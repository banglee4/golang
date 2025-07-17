package lppm

import "sippm/model/domain"

type UpdateDataRequest struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domain.Lppm {
	return domain.Lppm{
		Id:    request.Id,
		Nama:  request.Nama,
		Photo: request.Photo,
		Email: request.Email,
		Phone: request.Phone,
	}
}
