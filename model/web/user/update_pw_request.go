package user

import (
	"sippm/helper"
	"sippm/model/domain"
)

type UpdatePwRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func UpdatePwRequestToDomain(request UpdatePwRequest, userId int) domain.Users {
	return domain.Users{
		Id:       userId,
		Username: request.Username,
		Password: helper.HashedPassword(request.NewPassword),
	}
}
