package usulan

import (
	"sippm/model/domain"
	"sippm/model/web/file_laporan"
)

type CreateRequest struct {
	PengusulId     int                        `json:"pengusul_id"`
	Judul          string                     `json:"judul"`
	Fakultas       string                     `json:"fakultas"`
	Prodi          string                     `json:"prodi"`
	Jenis          string                     `json:"jenis"`
	Tahun          string                     `json:"tahun"`
	Status         string                     `json:"status"`
	Anggaran       string                     `json:"anggaran"`
	NamaAnggota1   *string                    `json:"nama_anggota_1"`
	NamaAnggota2   *string                    `json:"nama_anggota_2"`
	NamaAnggota3   *string                    `json:"nama_anggota_3"`
	NamaMahasiswa1 *string                    `json:"nama_mahasiswa_1"`
	NamaMahasiswa2 *string                    `json:"nama_mahasiswa_2"`
	NamaMahasiswa3 *string                    `json:"nama_mahasiswa_3"`
	NamaMahasiswa4 *string                    `json:"nama_mahasiswa_4"`
	FileLaporan    file_laporan.CreateRequest `json:"file_laporan"`
}

func CreateRequestToDomain(request CreateRequest) domain.Usulan {
	return domain.Usulan{
		PengusulId:     request.PengusulId,
		Judul:          request.Judul,
		Fakultas:       request.Fakultas,
		Prodi:          request.Prodi,
		Jenis:          request.Jenis,
		Tahun:          request.Tahun,
		Anggaran:       request.Anggaran,
		Status:         request.Status,
		NamaAnggota1:   request.NamaAnggota1,
		NamaAnggota2:   request.NamaAnggota2,
		NamaAnggota3:   request.NamaAnggota3,
		NamaMahasiswa1: request.NamaMahasiswa1,
		NamaMahasiswa2: request.NamaMahasiswa2,
		NamaMahasiswa3: request.NamaMahasiswa3,
		NamaMahasiswa4: request.NamaMahasiswa4,
		FileLaporan:    file_laporan.CreateRequestToDomain(request.FileLaporan),
	}
}
