package domain

type Users struct {
	Id       int    `gorm:"primary_key;column:id;auto_increment"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
	Dosen    Dosen  `gorm:"foreignKey:user_id;references:id"`
	Uppm     Uppm   `gorm:"foreignKey:user_id;references:id"`
	Lppm     Lppm   `gorm:"foreignKey:user_id;references:id"`
}

func (Users) TableName() string {
	return "users"
}
