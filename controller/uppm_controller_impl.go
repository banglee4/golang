package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/uppm"

	"github.com/gin-gonic/gin"
)

type UppmControllerImpl struct {
	Serv service.UppmService
}

func NewUppmControllerImpl(serv service.UppmService) *UppmControllerImpl {
	return &UppmControllerImpl{Serv: serv}
}

func (controller *UppmControllerImpl) UpdateData(context *gin.Context) {
	// Parse Request Body
	request := uppm.UpdateDataRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := controller.Serv.UpdateData(request)
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
