package user

import (
	"sippm/helper"
	"sippm/model/domain"
	"sippm/model/web/dosen"
)

type RegisterDosenRequest struct {
	Username string              `json:"username"`
	Password string              `json:"password"`
	Dosen    dosen.CreateRequest `json:"dosen"`
}

func RegisterDosenRequestToDomain(request RegisterDosenRequest) domain.Users {
	return domain.Users{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "Dosen",
		Dosen:    dosen.CreateRequestToDomain(request.Dosen),
	}
}
