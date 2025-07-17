package service

import (
	"sippm/model/web/user"
)

type UserService interface {
	RegisterAdmin(request user.RegisterAdminRequest) error
	RegisterDosen(request user.RegisterDosenRequest) error
	RegisterUppm(request user.RegisterUppmRequest) error
	RegisterLppm(request user.RegisterLppmRequest) error
	LoginAdmin(request user.LoginRequest) (*string, error)
	Login(request user.LoginRequest) (*string, error)
	UpdatePassword(request user.UpdatePwRequest) error
	ResetPassword(request user.ResetPwRequest) error
	GetDosenByUserId(userId int) (*user.DosenResponse, error)
	GetUppmByUserId(userId int) (*user.UppmResponse, error)
	GetLppmByUserId(userId int) (*user.LppmResponse, error)
	GetAllDosen() ([]user.DosenResponse, int, error)
	GetAllUppm() ([]user.UppmResponse, int, error)
	GetAllLppm() ([]user.LppmResponse, int, error)
	SearchDosenByName(name string) ([]user.DosenResponse, error)
	SearchUppmByName(name string) ([]user.UppmResponse, error)
	SearchLppmByName(name string) ([]user.LppmResponse, error)
	Delete(userId int) error
}
