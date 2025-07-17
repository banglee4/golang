package user

import (
	"sippm/helper"
	"sippm/model/domain"
)

type RegisterAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAdminRequestToDomain(request RegisterAdminRequest) domain.Users {
	return domain.Users{
		Username: request.Username,
		Password: helper.HashedPassword(request.Password),
		Role:     "Admin",
	}
}
