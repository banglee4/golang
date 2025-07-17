package repository

import "sippm/model/domain"

type UserRepository interface {
	Create(user domain.Users) error
	UpdatePassword(user domain.Users) error
	FindById(id int) (*domain.Users, error)
	FindByUsername(username string) (*domain.Users, error)
	FindDosenByUserId(userId int) (*domain.Users, error)
	FindUppmByUserId(userId int) (*domain.Users, error)
	FindLppmByUserId(userId int) (*domain.Users, error)
	GetAllDosen() ([]domain.Users, int, error)
	GetAllUppm() ([]domain.Users, int, error)
	GetAllLppm() ([]domain.Users, int, error)
	SearchDosenByName(name string) ([]domain.Users, error)
	SearchUppmByName(name string) ([]domain.Users, error)
	SearchLppmByName(name string) ([]domain.Users, error)
	DeleteById(id int) error
}
