package service

import (
	"gorm.io/gorm"
	"sippm/model/domain"
	"sippm/model/repository"
	"sippm/model/web/ketua"
)

type JabatanKetuaServiceImpl struct {
	Repo repository.JabatanKetuaRepository
	DB   *gorm.DB
}

func NewJabatanKetuaServiceImpl(repo repository.JabatanKetuaRepository, DB *gorm.DB) *JabatanKetuaServiceImpl {
	return &JabatanKetuaServiceImpl{Repo: repo, DB: DB}
}

func (serv *JabatanKetuaServiceImpl) Create(request ketua.CreateRequest) error {
	model := ketua.CreateReqToDomain(request)

	// Call Repo
	err := serv.Repo.Create(model, serv.DB)
	if err != nil {
		return err
	}

	return nil
}

func (serv *JabatanKetuaServiceImpl) Update(request ketua.UpdateRequest) error {
	model := ketua.UpdateReqToDomain(request)

	// Call Repo
	err := serv.Repo.Update(model, serv.DB)
	if err != nil {
		return err
	}

	return nil
}

func (serv *JabatanKetuaServiceImpl) GetAll() ([]domain.Ketua, error) {
	// Call Repo
	results, err := serv.Repo.GetAll(serv.DB)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (serv *JabatanKetuaServiceImpl) Delete(id int) error {
	// Call Repo
	err := serv.Repo.Delete(id, serv.DB)
	if err != nil {
		return err
	}

	return nil
}
