package usulan

import (
	"sippm/model/domain"
	"sippm/model/web/file_laporan"
	"sippm/model/web/user"
)

type Response struct {
	Id             int                   `json:"id"`
	Judul          string                `json:"judul"`
	Fakultas       string                `json:"fakultas"`
	Prodi          string                `json:"prodi"`
	Jenis          string                `json:"jenis"`
	Tahun          string                `json:"tahun"`
	Anggaran       string                `json:"anggaran"`
	Status         string                `json:"status"`
	NamaAnggota1   *string               `json:"nama_anggota_1"`
	NamaAnggota2   *string               `json:"nama_anggota_2"`
	NamaAnggota3   *string               `json:"nama_anggota_3"`
	NamaMahasiswa1 *string               `json:"nama_mahasiswa_1"`
	NamaMahasiswa2 *string               `json:"nama_mahasiswa_2"`
	NamaMahasiswa3 *string               `json:"nama_mahasiswa_3"`
	NamaMahasiswa4 *string               `json:"nama_mahasiswa_4"`
	Dosen          user.DosenResponse    `json:"dosen"`
	FileLaporan    file_laporan.Response `json:"file_laporan"`
}

func ToResponse(usulan domain.Usulan) *Response {
	return &Response{
		Id:             usulan.Id,
		Judul:          usulan.Judul,
		Fakultas:       usulan.Fakultas,
		Prodi:          usulan.Prodi,
		Jenis:          usulan.Jenis,
		Tahun:          usulan.Tahun,
		Anggaran:       usulan.Anggaran,
		Status:         usulan.Status,
		NamaAnggota1:   usulan.NamaAnggota1,
		NamaAnggota2:   usulan.NamaAnggota2,
		NamaAnggota3:   usulan.NamaAnggota3,
		NamaMahasiswa1: usulan.NamaMahasiswa1,
		NamaMahasiswa2: usulan.NamaMahasiswa2,
		NamaMahasiswa3: usulan.NamaMahasiswa3,
		NamaMahasiswa4: usulan.NamaMahasiswa4,
		Dosen:          *user.ToDosenResponse(usulan.User),
		FileLaporan:    *file_laporan.ToResponse(usulan.FileLaporan),
	}
}

func ToResponses(usulans []domain.Usulan) []Response {
	var results []Response

	for _, usulan := range usulans {
		results = append(results, *ToResponse(usulan))
	}

	return results
}
