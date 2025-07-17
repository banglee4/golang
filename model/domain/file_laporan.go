package domain

type FileLaporan struct {
	Id                     int    `gorm:"primary_key;column:id;auto_increment"`
	UsulanId               int    `gorm:"column:usulan_id"`
	FileUsulan             string `gorm:"column:file_usulan"`
	StatusUsulan           string `gorm:"column:status_usulan"`
	CatatanUsulan          string `gorm:"column:catatan_usulan"`
	FileLaporanKemajuan    string `gorm:"column:file_laporan_kemajuan"`
	StatusLaporanKemajuan  string `gorm:"column:status_laporan_kemajuan"`
	CatatanLaporanKemajuan string `gorm:"column:catatan_laporan_kemajuan"`
	FileLaporanAkhir       string `gorm:"column:file_laporan_akhir"`
	StatusLaporanAkhir     string `gorm:"column:status_laporan_akhir"`
	CatatanLaporanAkhir    string `gorm:"column:catatan_laporan_akhir"`
}

func (FileLaporan) TableName() string {
	return "filelaporan"
}
