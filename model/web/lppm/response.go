package lppm

import "sippm/model/domain"

type Response struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func ToResponse(lppm domain.Lppm) Response {
	return Response{
		Id:    lppm.Id,
		Nama:  lppm.Nama,
		Photo: lppm.Photo,
		Email: lppm.Email,
		Phone: lppm.Phone,
	}
}

func ToResponses(lppms []domain.Lppm) []Response {
	var results []Response

	for _, lppm := range lppms {
		results = append(results, ToResponse(lppm))
	}

	return results
}
