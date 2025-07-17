package usulan

import (
	"sippm/model/domain"
)

type SearchByJudulRequest struct {
	Judul string `json:"judul"`
}

func SearchByJudulRequestToDomain(request SearchByJudulRequest) domain.Usulan {
	return domain.Usulan{
		Judul: request.Judul,
	}
}
