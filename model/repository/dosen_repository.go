package repository

import "sippm/model/domain"

type DosenRepository interface {
	UpdateData(dosen domain.Dosen) error
}
