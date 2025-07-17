package domain

type Usulan struct {
	Id             int         `gorm:"primary_key;column:id;auto_increment"`
	PengusulId     int         `gorm:"column:pengusul_id"`
	Judul          string      `gorm:"column:judul"`
	Fakultas       string      `gorm:"column:fakultas"`
	Prodi          string      `gorm:"column:prodi"`
	Jenis          string      `gorm:"column:jenis_usulan"`
	Tahun          string      `gorm:"column:tahun_usulan"`
	Anggaran       string      `gorm:"column:anggaran"`
	Status         string      `gorm:"column:status"`
	NamaAnggota1   *string     `gorm:"column:nama_anggota_1"`
	NamaAnggota2   *string     `gorm:"column:nama_anggota_2"`
	NamaAnggota3   *string     `gorm:"column:nama_anggota_3"`
	NamaMahasiswa1 *string     `gorm:"column:nama_mahasiswa_1"`
	NamaMahasiswa2 *string     `gorm:"column:nama_mahasiswa_2"`
	NamaMahasiswa3 *string     `gorm:"column:nama_mahasiswa_3"`
	NamaMahasiswa4 *string     `gorm:"column:nama_mahasiswa_4"`
	User           Users       `gorm:"foreignKey:pengusul_id;references:id"`
	FileLaporan    FileLaporan `gorm:"foreignKey:usulan_id;references:id"`
}

func (Usulan) TableName() string {
	return "usulan"
}
