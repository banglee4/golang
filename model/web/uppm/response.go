package uppm

import "sippm/model/domain"

type Response struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Photo    string `json:"photo"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fakultas string `json:"fakultas"`
}

func ToResponse(uppm domain.Uppm) Response {
	return Response{
		Id:       uppm.Id,
		Nama:     uppm.Nama,
		Photo:    uppm.Photo,
		Email:    uppm.Email,
		Phone:    uppm.Phone,
		Fakultas: uppm.Fakultas,
	}
}
