package dosen

import "sippm/model/domain"

type Response struct {
	Id    int    `json:"id"`
	Nidn  string `json:"nidn"`
	Nama  string `json:"nama"`
	Photo string `json:"photo"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func ToResponse(dosen domain.Dosen) Response {
	return Response{
		Id:    dosen.Id,
		Nidn:  dosen.Nidn,
		Nama:  dosen.Nama,
		Photo: dosen.Photo,
		Email: dosen.Email,
		Phone: dosen.Phone,
	}
}

func ToResponses(dosens []domain.Dosen) []Response {
	var results []Response

	for _, dosen := range dosens {
		results = append(results, ToResponse(dosen))
	}

	return results
}
