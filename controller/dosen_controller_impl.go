package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/dosen"

	"github.com/gin-gonic/gin"
)

type DosenControllerImpl struct {
	Service service.DosenService
}

func NewDosenControllerImpl(service service.DosenService) *DosenControllerImpl {
	return &DosenControllerImpl{Service: service}
}

func (controller *DosenControllerImpl) UpdateData(context *gin.Context) {
	// Parse Request Body
	request := dosen.UpdateDataRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := controller.Service.UpdateData(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}
