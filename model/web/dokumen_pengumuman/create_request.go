package dokumen_pengumuman

import "sippm/model/domain"

type CreateRequest struct {
	NamaDokumen string `json:"nama_dokumen"`
	Url         string `json:"url"`
}

func CreateReqToDomain(request CreateRequest) domain.DokumenPengumuman {
	return domain.DokumenPengumuman{
		NamaDokumen: request.NamaDokumen,
		Url:         request.Url,
	}
}

func SomeCreateRequestToDomains(requests []CreateRequest) []domain.DokumenPengumuman {
	var results []domain.DokumenPengumuman

	for _, request := range requests {
		results = append(results, CreateReqToDomain(request))
	}

	return results
}
