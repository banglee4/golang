package service

import (
	"fmt"
	"sippm/model/repository"
	"sippm/model/web/usulan"
)

type UsulanServiceImpl struct {
	Repo repository.UsulanRepository
}

func NewUsulanServiceImpl(repo repository.UsulanRepository) *UsulanServiceImpl {
	return &UsulanServiceImpl{Repo: repo}
}

func (serv *UsulanServiceImpl) Create(request usulan.CreateRequest) error {
	model := usulan.CreateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Create(model)
	if err != nil {
		return fmt.Errorf("judul already exists")
	}

	return nil
}

func (serv *UsulanServiceImpl) Update(request usulan.UpdateRequest) error {
	model := usulan.UpdateRequestToDomain(request)

	// Call Repo
	err := serv.Repo.Update(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *UsulanServiceImpl) GetByJenisAndStatus(request usulan.JenisStatusRequest) ([]usulan.Response, int, error) {
	// Call Repo
	results, total, err := serv.Repo.GetByJenisAndStatus(usulan.JenisStatusRequestToDomain(request))
	if err != nil {
		return nil, 0, err
	}

	return usulan.ToResponses(results), total, nil
}

func (serv *UsulanServiceImpl) SearchByJudul(request usulan.SearchByJudulRequest) ([]usulan.Response, int, error) {
	// Call Repo
	results, total, err := serv.Repo.SearchByJudul(usulan.SearchByJudulRequestToDomain(request))
	if err != nil {
		return nil, 0, err
	}

	return usulan.ToResponses(results), total, nil
}

func (serv *UsulanServiceImpl) GetAll(jenis string, fakultas string) ([]usulan.Response, int, error) {
	// Call Repo
	results, total, err := serv.Repo.GetAll(jenis, fakultas)
	if err != nil {
		return nil, 0, err
	}

	return usulan.ToResponses(results), total, nil
}

func (serv *UsulanServiceImpl) GetAllValidated(jenis string) ([]usulan.Response, int, error) {
	// Call Repo
	results, total, err := serv.Repo.GetAllValidated(jenis)
	if err != nil {
		return nil, 0, err
	}

	return usulan.ToResponses(results), total, nil
}

func (serv *UsulanServiceImpl) GetAllByPengusulId(request usulan.GetAllByPengusul) ([]usulan.Response, int, error) {
	// Call Repo
	results, total, err := serv.Repo.GetAllByPengusulId(request.PengusulId, request.Jenis)
	if err != nil {
		return nil, 0, err
	}

	return usulan.ToResponses(results), total, nil
}

func (serv *UsulanServiceImpl) Filter(request usulan.FilterRequest) ([]usulan.Response, error) {
	// Call Repo
	results, err := serv.Repo.Filter(usulan.FilterRequestToDomain(request), request.TahunMulai, request.TahunAkhir)
	if err != nil {
		return nil, err
	}

	return usulan.ToResponses(results), nil
}

func (serv *UsulanServiceImpl) UpdateStatus(request usulan.UpdateStatusRequest) error {
	model := usulan.UpdateStatusRequestToDomain(request)

	// Call Repo
	err := serv.Repo.UpdateStatus(model)
	if err != nil {
		return err
	}

	return nil
}

func (serv *UsulanServiceImpl) Delete(id int) error {
	// Call Repo
	err := serv.Repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
