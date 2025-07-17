package domain

type Pengumuman struct {
	Id            int                 `gorm:"primary_key;column:id;auto_increment"`
	Judul         string              `gorm:"column:judul"`
	TanggalDibuat string              `gorm:"column:created_at"`
	Dokumen       []DokumenPengumuman `gorm:"foreignKey:pengumuman_id;references:id"`
}

type DokumenPengumuman struct {
	Id           int    `gorm:"primary_key;column:id;auto_increment"`
	PengumumanId int    `gorm:"column:pengumuman_id"`
	NamaDokumen  string `gorm:"column:nama_dokument"`
	Url          string `gorm:"column:url"`
}

func (Pengumuman) TableName() string {
	return "pengumuman"
}

func (DokumenPengumuman) TableName() string {
	return "dokumenpengumuman"
}
