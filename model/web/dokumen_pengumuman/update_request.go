package dokumen_pengumuman

import "sippm/model/domain"

type UpdateRequest struct {
	Id          int    `json:"id"`
	NamaDokumen string `json:"nama_dokumen"`
	Url         string `json:"url"`
}

func UpdateRequestToDomain(request UpdateRequest) domain.DokumenPengumuman {
	return domain.DokumenPengumuman{
		Id:          request.Id,
		NamaDokumen: request.NamaDokumen,
		Url:         request.Url,
	}
}

func SomeUpdateRequestToDomains(requests []UpdateRequest) []domain.DokumenPengumuman {
	var results []domain.DokumenPengumuman

	for _, request := range requests {
		results = append(results, UpdateRequestToDomain(request))
	}

	return results
}
