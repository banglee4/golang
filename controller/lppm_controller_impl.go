package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/lppm"

	"github.com/gin-gonic/gin"
)

type LppmControllerImpl struct {
	Serv service.LppmService
}

func NewLppmControllerImpl(serv service.LppmService) *LppmControllerImpl {
	return &LppmControllerImpl{Serv: serv}
}

func (controller *LppmControllerImpl) UpdateData(context *gin.Context) {
	// Parse Request Body
	request := lppm.UpdateDataRequest{}
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
