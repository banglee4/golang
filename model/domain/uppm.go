package domain

type Uppm struct {
	Id       int    `gorm:"primary_key;column:id;auto_increment"`
	UserId   int    `gorm:"column:user_id"`
	Nama     string `gorm:"column:nama"`
	Photo    string `gorm:"column:photo"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
	Fakultas string `gorm:"column:fakultas"`
}

func (Uppm) TableName() string {
	return "uppm"
}
