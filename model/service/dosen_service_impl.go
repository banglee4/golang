package service

import (
	"sippm/model/repository"
	"sippm/model/web/dosen"
)

type DosenServiceImpl struct {
	Repo repository.DosenRepository
}

func NewDosenServiceImpl(repo repository.DosenRepository) *DosenServiceImpl {
	return &DosenServiceImpl{Repo: repo}
}

func (serv *DosenServiceImpl) UpdateData(request dosen.UpdateDataRequest) error {
	model := dosen.UpdateDataRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateData(model)
	if err != nil {
		return err
	}

	return nil
}
