package service

import (
	"sippm/model/repository"
	"sippm/model/web/uppm"
)

type UppmServiceImpl struct {
	Repo repository.UppmRepository
}

func NewUppmServiceImpl(repo repository.UppmRepository) *UppmServiceImpl {
	return &UppmServiceImpl{Repo: repo}
}

func (serv UppmServiceImpl) UpdateData(request uppm.UpdateDataRequest) error {
	model := uppm.UpdateDataRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateData(model)
	if err != nil {
		return err
	}

	return nil
}
