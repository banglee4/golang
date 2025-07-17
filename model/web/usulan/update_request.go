package usulan

import (
	"sippm/model/domain"
)

type UpdateRequest struct {
	Id             int     `json:"id"`
	Judul          string  `json:"judul"`
	Fakultas       string  `json:"fakultas"`
	Prodi          string  `json:"prodi"`
	Jenis          string  `json:"jenis"`
	Anggaran       string  `json:"anggaran"`
	Tahun          string  `json:"tahun"`
	NamaAnggota1   *string `json:"nama_anggota1"`
	NamaAnggota2   *string `json:"nama_anggota2"`
	NamaAnggota3   *string `json:"nama_anggota3"`
	NamaMahasiswa1 *string `json:"nama_mahasiswa1"`
	NamaMahasiswa2 *string `json:"nama_mahasiswa2"`
	NamaMahasiswa3 *string `json:"nama_mahasiswa3"`
	NamaMahasiswa4 *string `json:"nama_mahasiswa4"`
}

func UpdateRequestToDomain(request UpdateRequest) domain.Usulan {
	return domain.Usulan{
		Id:             request.Id,
		Judul:          request.Judul,
		Fakultas:       request.Fakultas,
		Prodi:          request.Prodi,
		Jenis:          request.Jenis,
		Anggaran:       request.Anggaran,
		Tahun:          request.Tahun,
		NamaAnggota1:   request.NamaAnggota1,
		NamaAnggota2:   request.NamaAnggota2,
		NamaAnggota3:   request.NamaAnggota3,
		NamaMahasiswa1: request.NamaMahasiswa1,
		NamaMahasiswa2: request.NamaMahasiswa2,
		NamaMahasiswa3: request.NamaMahasiswa3,
		NamaMahasiswa4: request.NamaMahasiswa4,
	}
}
