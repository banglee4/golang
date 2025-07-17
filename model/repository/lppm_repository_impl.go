package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type LppmRepositoryImpl struct {
	DB *gorm.DB
}

func NewLppmRepositoryImpl(DB *gorm.DB) *LppmRepositoryImpl {
	return &LppmRepositoryImpl{DB: DB}
}

func (repo *LppmRepositoryImpl) UpdateData(lppm domain.Lppm) error {
	err := repo.DB.
		Model(&lppm).
		Where("id = ?", lppm.Id).
		Updates(map[string]interface{}{
			"nama":  lppm.Nama,
			"photo": lppm.Photo,
			"email": lppm.Email,
			"phone": lppm.Phone,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update data")
	}

	return nil
}
