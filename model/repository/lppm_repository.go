package repository

import "sippm/model/domain"

type LppmRepository interface {
	UpdateData(lppm domain.Lppm) error
}
