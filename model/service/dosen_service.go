package service

import "sippm/model/web/dosen"

type DosenService interface {
	UpdateData(request dosen.UpdateDataRequest) error
}
