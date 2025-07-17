package user

import (
	"sippm/helper"
	"sippm/model/domain"
)

type ResetPwRequest struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
}

func ResetPwRequestToDomain(request ResetPwRequest) domain.Users {
	return domain.Users{
		Id:       request.UserId,
		Password: helper.HashedPassword(request.Password),
	}
}
