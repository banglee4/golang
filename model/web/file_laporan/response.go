package file_laporan

import "sippm/model/domain"

type Response struct {
	Id                     int    `json:"id"`
	FileUsulan             string `json:"file_usulan"`
	StatusUsulan           string `json:"status_usulan"`
	CatatanUsulan          string `json:"catatan_usulan"`
	FileLaporanKemajuan    string `json:"file_laporan_kemajuan"`
	StatusLaporanKemajuan  string `json:"status_laporan_kemajuan"`
	CatatanLaporanKemajuan string `json:"catatan_laporan_kemajuan"`
	FileLaporanAkhir       string `json:"file_laporan_akhir"`
	StatusLaporanAkhir     string `json:"status_laporan_akhir"`
	CatatanLaporanAkhir    string `json:"catatan_laporan_akhir"`
}

func ToResponse(file domain.FileLaporan) *Response {
	return &Response{
		Id:                     file.Id,
		FileUsulan:             file.FileUsulan,
		StatusUsulan:           file.StatusUsulan,
		CatatanUsulan:          file.CatatanUsulan,
		FileLaporanKemajuan:    file.FileLaporanKemajuan,
		StatusLaporanKemajuan:  file.StatusLaporanKemajuan,
		CatatanLaporanKemajuan: file.CatatanLaporanKemajuan,
		FileLaporanAkhir:       file.FileLaporanAkhir,
		StatusLaporanAkhir:     file.StatusLaporanAkhir,
		CatatanLaporanAkhir:    file.CatatanLaporanAkhir,
	}
}
