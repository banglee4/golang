package repository

import "sippm/model/domain"

type UppmRepository interface {
	UpdateData(uppm domain.Uppm) error
}
