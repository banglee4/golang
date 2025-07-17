package repository

import (
	"fmt"
	"sippm/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (repo *UserRepositoryImpl) Create(user domain.Users) error {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (repo *UserRepositoryImpl) UpdatePassword(user domain.Users) error {
	err := repo.DB.
		Model(&user).
		Where("id = ?", user.Id).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}

func (repo *UserRepositoryImpl) FindById(id int) (*domain.Users, error) {
	var user domain.Users
	err := repo.DB.
		Select("id", "username", "role").
		Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindByUsername(username string) (*domain.Users, error) {
	var user domain.Users
	err := repo.DB.
		Select("id", "username", "password", "role").
		Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindDosenByUserId(userId int) (*domain.Users, error) {
	var user domain.Users
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Dosen").
		Where("id = ? AND role = 'Dosen'", userId).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get dosen by id: %w", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindUppmByUserId(userId int) (*domain.Users, error) {
	var user domain.Users
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Uppm").
		Where("id = ? AND role = 'UPPM'", userId).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get uppm by id: %w", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindLppmByUserId(userId int) (*domain.Users, error) {
	var user domain.Users
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Lppm").
		Where("id = ? AND role = 'LPPM'", userId).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get lppm by id: %w", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetAllDosen() ([]domain.Users, int, error) {
	var users []domain.Users
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Users{}).Where("role = ?", "Dosen").Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count dosen: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Dosen").
		Where("role = ?", "Dosen").
		Order("created_at DESC").
		Find(&users).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all dosen: %w", err)
	}

	return users, int(total), nil
}

func (repo *UserRepositoryImpl) GetAllUppm() ([]domain.Users, int, error) {
	var users []domain.Users
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Users{}).Where("role = ?", "UPPM").Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count uppm: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Uppm").
		Where("role = ?", "UPPM").
		Order("created_at DESC").
		Find(&users).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all uppm: %w", err)
	}

	return users, int(total), nil
}

func (repo *UserRepositoryImpl) GetAllLppm() ([]domain.Users, int, error) {
	var users []domain.Users
	var total int64

	// Hitung total data dosen
	if err := repo.DB.Model(&domain.Users{}).Where("role = ?", "LPPM").Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count lppm: %w", err)
	}

	// Ambil data dosen dengan pagination
	err := repo.DB.
		Select("id", "username", "role").
		Preload("Lppm").
		Where("role = ?", "LPPM").
		Order("created_at DESC").
		Find(&users).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all lppm: %w", err)
	}

	return users, int(total), nil
}

func (repo *UserRepositoryImpl) SearchDosenByName(name string) ([]domain.Users, error) {
	var results []domain.Users

	query := repo.DB.
		Select("users.id, users.username, users.role").
		Joins("JOIN dosen ON users.id = dosen.user_id").
		Where("users.role = ?", "Dosen")

	if name != "" {
		query = query.Where("dosen.nama LIKE ?", "%"+name+"%")
	}

	err := query.Preload("Dosen").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *UserRepositoryImpl) SearchUppmByName(name string) ([]domain.Users, error) {
	var results []domain.Users

	query := repo.DB.
		Select("users.id, users.username, users.role").
		Joins("JOIN uppm ON users.id = uppm.user_id").
		Where("users.role = ?", "UPPM")

	if name != "" {
		query = query.Where("uppm.nama LIKE ?", "%"+name+"%")
	}

	err := query.Preload("Uppm").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *UserRepositoryImpl) SearchLppmByName(name string) ([]domain.Users, error) {
	var results []domain.Users

	query := repo.DB.
		Select("users.id, users.username, users.role").
		Joins("JOIN lppm ON users.id = lppm.user_id").
		Where("users.role = ?", "LPPM")

	if name != "" {
		query = query.Where("lppm.nama LIKE ?", "%"+name+"%")
	}

	err := query.Preload("Lppm").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *UserRepositoryImpl) DeleteById(id int) error {
	tx := repo.DB.Begin()

	// Cek apakah user ada
	var user domain.Users
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("user not found: %w", err)
	}

	// Hapus entitas terkait berdasarkan role
	switch user.Role {
	case "Dosen":
		if err := tx.Where("user_id = ?", id).Delete(&domain.Dosen{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete related dosen: %w", err)
		}
	case "UPPM":
		if err := tx.Where("user_id = ?", id).Delete(&domain.Uppm{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete related uppm: %w", err)
		}
	case "LPPM":
		if err := tx.Where("user_id = ?", id).Delete(&domain.Lppm{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete related lppm: %w", err)
		}
	}

	// Periksa apakah user memiliki usulan
	var count int64
	tx.Model(&domain.Usulan{}).Where("pengusul_id = ?", id).Count(&count)
	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("cannot delete user: user has existing usulan references")
	}

	// Hapus user
	if err := tx.Delete(&domain.Users{}, id).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return tx.Commit().Error
}
