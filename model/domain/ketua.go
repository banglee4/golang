package domain

type Ketua struct {
	Id        int    `gorm:"primary_key;column:id;auto_increment"`
	NamaDosen string `gorm:"column:nama_dosen"`
	Jabatan   string `gorm:"column:jabatan"`
}

func (Ketua) TableName() string {
	return "jabatanketua"
}
