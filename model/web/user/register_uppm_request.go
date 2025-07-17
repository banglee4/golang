package user

import (
	"sippm/helper"
	"sippm/model/domain"
	"sippm/model/web/uppm"
)

type RegisterUppmRequest struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	Uppm     uppm.CreateRequest `json:"uppm"`
}

func RegisterUppmRequestToDomain(request RegisterUppmRequest) domain.Users {
	return domain.Users{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "UPPM",
		Uppm:     uppm.CreateRequestToDomain(request.Uppm),
	}
}
