package usulan

import "sippm/model/domain"

type FilterRequest struct {
	Prodi      string `json:"prodi"`
	Jenis      string `json:"jenis_usulan"`
	TahunMulai string `json:"tahun_usulan"`
	TahunAkhir string `json:"tahun_akhir"`
}

func FilterRequestToDomain(request FilterRequest) domain.Usulan {
	return domain.Usulan{
		Prodi: request.Prodi,
		Jenis: request.Jenis,
	}
}
