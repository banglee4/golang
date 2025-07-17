package service

import (
	"sippm/model/web/usulan"
)

type UsulanService interface {
	Create(request usulan.CreateRequest) error
	Update(request usulan.UpdateRequest) error
	GetByJenisAndStatus(request usulan.JenisStatusRequest) ([]usulan.Response, int, error)
	SearchByJudul(request usulan.SearchByJudulRequest) ([]usulan.Response, int, error)
	Filter(request usulan.FilterRequest) ([]usulan.Response, error)
	GetAll(jenis string, fakultas string) ([]usulan.Response, int, error)
	GetAllValidated(jenis string) ([]usulan.Response, int, error)
	GetAllByPengusulId(request usulan.GetAllByPengusul) ([]usulan.Response, int, error)
	UpdateStatus(request usulan.UpdateStatusRequest) error
	Delete(id int) error
}
