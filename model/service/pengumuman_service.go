package service

import "sippm/model/web/pengumuman"

type PengumumanService interface {
	Create(request pengumuman.CreateRequest) error
	Update(request pengumuman.UpdateRequest) error
	GetAll() ([]pengumuman.Response, error)
	Delete(id int) error
}
