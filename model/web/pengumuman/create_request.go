package pengumuman

import (
	"sippm/model/domain"
	"sippm/model/web/dokumen_pengumuman"
)

type CreateRequest struct {
	Judul   string                             `json:"judul"`
	Dokumen []dokumen_pengumuman.CreateRequest `json:"dokumen"`
	Dibuat  string                             `json:"dibuat_tanggal"`
}

func CreateRequestToDomain(request CreateRequest) domain.Pengumuman {
	return domain.Pengumuman{
		Judul:         request.Judul,
		Dokumen:       dokumen_pengumuman.SomeCreateRequestToDomains(request.Dokumen),
		TanggalDibuat: request.Dibuat,
	}
}
