package pengumuman

import (
	"sippm/model/domain"
	"sippm/model/web/dokumen_pengumuman"
)

type Response struct {
	Id            int                           `json:"id"`
	Judul         string                        `json:"judul"`
	Dokumen       []dokumen_pengumuman.Response `json:"dokumen"`
	TanggalDibuat string                        `json:"tanggal_dibuat"`
}

func ToResponse(model domain.Pengumuman) *Response {
	return &Response{
		Id:            model.Id,
		Judul:         model.Judul,
		Dokumen:       dokumen_pengumuman.ToResponses(model.Dokumen),
		TanggalDibuat: model.TanggalDibuat,
	}
}

func ToResponses(models []domain.Pengumuman) []Response {
	var responses []Response

	for _, model := range models {
		responses = append(responses, *ToResponse(model))
	}

	return responses
}
