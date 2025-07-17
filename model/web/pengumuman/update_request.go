package pengumuman

import (
	"sippm/model/domain"
	"sippm/model/web/dokumen_pengumuman"
)

type UpdateRequest struct {
	Id      int                                `json:"id"`
	Judul   string                             `json:"judul"`
	Dokumen []dokumen_pengumuman.UpdateRequest `json:"dokumen"`
	Dibuat  string                             `json:"dibuat_tanggal"`
}

func UpdateRequestToDomain(request UpdateRequest) domain.Pengumuman {
	return domain.Pengumuman{
		Id:            request.Id,
		Judul:         request.Judul,
		Dokumen:       dokumen_pengumuman.SomeUpdateRequestToDomains(request.Dokumen),
		TanggalDibuat: request.Dibuat,
	}
}
