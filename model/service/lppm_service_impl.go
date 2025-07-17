package service

import (
	"sippm/model/repository"
	"sippm/model/web/lppm"
)

type LppmServiceImpl struct {
	Repo repository.LppmRepository
}

func NewLppmServiceImpl(repo repository.LppmRepository) *LppmServiceImpl {
	return &LppmServiceImpl{Repo: repo}
}

func (serv *LppmServiceImpl) UpdateData(request lppm.UpdateDataRequest) error {
	model := lppm.UpdateDataRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateData(model)
	if err != nil {
		return err
	}

	return nil
}
