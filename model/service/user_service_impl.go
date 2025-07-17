package service

import (
	"fmt"
	"sippm/helper"
	"sippm/model/repository"
	"sippm/model/web/user"
)

type UserServiceImpl struct {
	Repo      repository.UserRepository
	SecretKey string
}

func NewUserServiceImpl(repo repository.UserRepository, secretKey string) *UserServiceImpl {
	return &UserServiceImpl{Repo: repo, SecretKey: secretKey}
}

func (service *UserServiceImpl) RegisterAdmin(request user.RegisterAdminRequest) error {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result != nil {
		return fmt.Errorf("user already exists")
	}

	model := user.RegisterAdminRequestToDomain(request)
	err := service.Repo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) RegisterDosen(request user.RegisterDosenRequest) error {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result != nil {
		return fmt.Errorf("user already exists")
	}

	model := user.RegisterDosenRequestToDomain(request)
	err := service.Repo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) RegisterUppm(request user.RegisterUppmRequest) error {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result != nil {
		return fmt.Errorf("user already exists")
	}

	model := user.RegisterUppmRequestToDomain(request)
	err := service.Repo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) RegisterLppm(request user.RegisterLppmRequest) error {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result != nil {
		return fmt.Errorf("user already exists")
	}

	model := user.RegisterLppmRequestToDomain(request)
	err := service.Repo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) LoginAdmin(request user.LoginRequest) (*string, error) {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check Username
	if !helper.CheckPassword(result.Password, request.Password) {
		return nil, fmt.Errorf("password incorrect")
	}

	// Check Role
	if result.Role != "Admin" {
		return nil, fmt.Errorf("user not authorized")
	}

	// Generate Token
	token, errToken := helper.GenerateJWT(result.Id, result.Role, service.SecretKey)
	if errToken != nil {
		return nil, errToken
	}

	return &token, nil
}

func (service *UserServiceImpl) Login(request user.LoginRequest) (*string, error) {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check Username
	if !helper.CheckPassword(result.Password, request.Password) {
		return nil, fmt.Errorf("password incorrect")
	}

	// Check Role
	if result.Role == "Admin" {
		return nil, fmt.Errorf("user not authorized")
	}

	// Generate Token
	token, errToken := helper.GenerateJWT(result.Id, result.Role, service.SecretKey)
	if errToken != nil {
		return nil, errToken
	}

	return &token, nil
}

func (service *UserServiceImpl) ResetPassword(request user.ResetPwRequest) error {
	model := user.ResetPwRequestToDomain(request)

	// Call Repo
	err := service.Repo.UpdatePassword(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) UpdatePassword(request user.UpdatePwRequest) error {
	// Call Repo
	result, _ := service.Repo.FindByUsername(request.Username)
	if result == nil {
		return fmt.Errorf("user not found")
	}

	// Check Password
	if !helper.CheckPassword(result.Password, request.OldPassword) {
		return fmt.Errorf("password incorrect")
	}

	model := user.UpdatePwRequestToDomain(request, result.Id)
	err := service.Repo.UpdatePassword(model)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) GetDosenByUserId(userId int) (*user.DosenResponse, error) {
	// Call Repo
	result, err := service.Repo.FindDosenByUserId(userId)
	if err != nil {
		return nil, err
	}

	return user.ToDosenResponse(*result), nil
}

func (service *UserServiceImpl) GetUppmByUserId(userId int) (*user.UppmResponse, error) {
	// Call Repo
	result, err := service.Repo.FindUppmByUserId(userId)
	if err != nil {
		return nil, err
	}

	return user.ToUppmResponse(*result), nil
}

func (service *UserServiceImpl) GetLppmByUserId(userId int) (*user.LppmResponse, error) {
	// Call Repo
	result, err := service.Repo.FindLppmByUserId(userId)
	if err != nil {
		return nil, err
	}

	return user.ToLppmResponse(*result), nil
}

func (service *UserServiceImpl) GetAllDosen() ([]user.DosenResponse, int, error) {
	// Call Repo
	results, totalData, err := service.Repo.GetAllDosen()
	if err != nil {
		return nil, 0, err
	}

	return user.ToDosenResponses(results), totalData, nil
}

func (service *UserServiceImpl) GetAllUppm() ([]user.UppmResponse, int, error) {
	// Call Repo
	results, totalData, err := service.Repo.GetAllUppm()
	if err != nil {
		return nil, 0, err
	}

	return user.ToUppmResponses(results), totalData, nil
}

func (service *UserServiceImpl) GetAllLppm() ([]user.LppmResponse, int, error) {
	// Call Repo
	results, totalData, err := service.Repo.GetAllLppm()
	if err != nil {
		return nil, 0, err
	}

	return user.ToLppmResponses(results), totalData, nil
}

func (service *UserServiceImpl) SearchDosenByName(name string) ([]user.DosenResponse, error) {
	// Call Repo
	results, err := service.Repo.SearchDosenByName(name)
	if err != nil {
		return nil, err
	}

	return user.ToDosenResponses(results), nil
}

func (service *UserServiceImpl) SearchUppmByName(name string) ([]user.UppmResponse, error) {
	// Call Repo
	results, err := service.Repo.SearchUppmByName(name)
	if err != nil {
		return nil, err
	}

	return user.ToUppmResponses(results), nil
}

func (service *UserServiceImpl) SearchLppmByName(name string) ([]user.LppmResponse, error) {
	// Call Repo
	results, err := service.Repo.SearchLppmByName(name)
	if err != nil {
		return nil, err
	}

	return user.ToLppmResponses(results), nil
}

func (service *UserServiceImpl) Delete(userId int) error {
	// Call Repo
	err := service.Repo.DeleteById(userId)
	if err != nil {
		return err
	}

	return nil
}
