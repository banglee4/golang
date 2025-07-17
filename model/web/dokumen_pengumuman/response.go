package dokumen_pengumuman

import "sippm/model/domain"

type Response struct {
	Id          int    `json:"id"`
	NamaDokumen string `json:"nama_dokumen"`
	Url         string `json:"url"`
}

func ToResponse(model domain.DokumenPengumuman) *Response {
	return &Response{
		Id:          model.Id,
		NamaDokumen: model.NamaDokumen,
		Url:         model.Url,
	}
}

func ToResponses(model []domain.DokumenPengumuman) []Response {
	var responses []Response

	for _, dokumen := range model {
		responses = append(responses, *ToResponse(dokumen))
	}

	return responses
}
