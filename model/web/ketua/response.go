package ketua

import "sippm/model/domain"

type Response struct {
	Id        int    `json:"id"`
	NamaDosen string `json:"nama_dosen"`
	Jabatan   string `json:"jabatan"`
}

func ToResponse(ketua domain.Ketua) *Response {
	return &Response{
		Id:        ketua.Id,
		NamaDosen: ketua.NamaDosen,
		Jabatan:   ketua.Jabatan,
	}
}

func ToResponses(ketuas []domain.Ketua) []Response {
	var responses []Response

	for _, ketua := range ketuas {
		responses = append(responses, *ToResponse(ketua))
	}

	return responses
}
