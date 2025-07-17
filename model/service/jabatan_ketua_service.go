package service

import (
	"sippm/model/domain"
	"sippm/model/web/ketua"
)

type JabatanKetuaService interface {
	Create(request ketua.CreateRequest) error
	Update(request ketua.UpdateRequest) error
	GetAll() ([]domain.Ketua, error)
	Delete(id int) error
}
