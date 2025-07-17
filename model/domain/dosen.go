package domain

type Dosen struct {
	Id     int    `gorm:"primary_key;column:id;auto_increment"`
	UserId int    `gorm:"column:user_id"`
	Nidn   string `gorm:"column:nidn"`
	Nama   string `gorm:"column:nama"`
	Photo  string `gorm:"column:photo"`
	Email  string `gorm:"column:email"`
	Phone  string `gorm:"column:phone"`
}

func (Dosen) TableName() string {
	return "dosen"
}
