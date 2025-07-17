package usulan

import "sippm/model/domain"

type UpdateStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func UpdateStatusRequestToDomain(request UpdateStatusRequest) domain.Usulan {
	return domain.Usulan{
		Id:     request.Id,
		Status: request.Status,
	}
}
