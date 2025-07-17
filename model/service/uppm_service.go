package service

import (
	"sippm/model/web/uppm"
)

type UppmService interface {
	UpdateData(request uppm.UpdateDataRequest) error
}
