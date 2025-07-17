package repository

import "sippm/model/domain"

type UsulanRepository interface {
	Create(usulan domain.Usulan) error
	Update(usulan domain.Usulan) error
	GetByJenisAndStatus(usulan domain.Usulan) ([]domain.Usulan, int, error)
	SearchByJudul(usulan domain.Usulan) ([]domain.Usulan, int, error)
	Filter(usulan domain.Usulan, tahunMulai, tahunAkhir string) ([]domain.Usulan, error)
	GetAll(jenis string, fakultas string) ([]domain.Usulan, int, error)
	GetAllValidated(jenis string) ([]domain.Usulan, int, error)
	GetAllByPengusulId(pengusulId int, jenis string) ([]domain.Usulan, int, error)
	UpdateStatus(usulan domain.Usulan) error
	Delete(id int) error
}
