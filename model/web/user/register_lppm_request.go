package user

import (
	"sippm/helper"
	"sippm/model/domain"
	"sippm/model/web/lppm"
)

type RegisterLppmRequest struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	Lppm     lppm.CreateRequest `json:"lppm"`
}

func RegisterLppmRequestToDomain(request RegisterLppmRequest) domain.Users {
	return domain.Users{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "LPPM",
		Lppm:     lppm.CreateRequestToDomain(request.Lppm),
	}
}
