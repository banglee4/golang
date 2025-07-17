package service

import (
	"sippm/model/web/lppm"
)

type LppmService interface {
	UpdateData(request lppm.UpdateDataRequest) error
}
