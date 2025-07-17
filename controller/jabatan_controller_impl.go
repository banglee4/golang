package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/ketua"
	"strconv"
)

type JabatanControllerImpl struct {
	Serv service.JabatanKetuaService
}

func NewJabatanControllerImpl(serv service.JabatanKetuaService) *JabatanControllerImpl {
	return &JabatanControllerImpl{Serv: serv}
}

func (cont JabatanControllerImpl) Create(context *gin.Context) {
	// Parse Request Body
	request := ketua.CreateRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.Create(request)
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

func (cont JabatanControllerImpl) Update(context *gin.Context) {
	// Parse Request Body
	request := ketua.UpdateRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.Update(request)
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

func (cont JabatanControllerImpl) GetAll(context *gin.Context) {
	// Call Service
	response, errServ := cont.Serv.GetAll()
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ketua.ToResponses(response),
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont JabatanControllerImpl) Delete(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	errServ := cont.Serv.Delete(idFinal)
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
