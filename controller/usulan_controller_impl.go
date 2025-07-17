package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/pagination"
	"sippm/model/web/usulan"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsulanControllerImpl struct {
	Service service.UsulanService
}

func NewUsulanControllerImpl(service service.UsulanService) *UsulanControllerImpl {
	return &UsulanControllerImpl{Service: service}
}

func (cont *UsulanControllerImpl) Create(context *gin.Context) {
	// Parse Request Body
	request := usulan.CreateRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.Create(request)
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

func (cont *UsulanControllerImpl) Update(context *gin.Context) {
	// Parse Request Body
	request := usulan.UpdateRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.Update(request)
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

func (cont *UsulanControllerImpl) GetByJenisAndStatus(context *gin.Context) {
	// Parse Request Body
	request := usulan.JenisStatusRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	results, totalData, errServ := cont.Service.GetByJenisAndStatus(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	data, errData := pagination.ResponseToJson(pagination.Response{
		Data:      results,
		TotalData: totalData,
	})
	if errData != nil {
		exception.ErrorHandler(context, errData)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) SearchByJudul(context *gin.Context) {
	// Parse Request Body
	request := usulan.SearchByJudulRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	results, totalData, errServ := cont.Service.SearchByJudul(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	data, errData := pagination.ResponseToJson(pagination.Response{
		Data:      results,
		TotalData: totalData,
	})
	if errData != nil {
		exception.ErrorHandler(context, errData)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) Filter(context *gin.Context) {
	// Parse Request Body
	request := usulan.FilterRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	results, errServ := cont.Service.Filter(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   results,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) UpdateStatus(context *gin.Context) {
	// Parse Request Body
	request := usulan.UpdateStatusRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateStatus(request)
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

func (cont *UsulanControllerImpl) GetAll(context *gin.Context) {
	// Call Service
	results, totalData, errServ := cont.Service.GetAll(context.Query("jenis"), context.Query("fakultas"))
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	data, errData := pagination.ResponseToJson(pagination.Response{
		Data:      results,
		TotalData: totalData,
	})
	if errData != nil {
		exception.ErrorHandler(context, errData)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) GetAllValidated(context *gin.Context) {
	// Call Service
	results, totalData, errServ := cont.Service.GetAllValidated(context.Query("jenis"))
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	data, errData := pagination.ResponseToJson(pagination.Response{
		Data:      results,
		TotalData: totalData,
	})
	if errData != nil {
		exception.ErrorHandler(context, errData)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) GetAllByPengusulId(context *gin.Context) {
	// Parse Request Body
	pengusulId := context.Query("pengusul_id")
	idFinal, _ := strconv.Atoi(pengusulId)
	jenis := context.Query("jenis")

	request := usulan.GetAllByPengusul{
		PengusulId: idFinal,
		Jenis:      jenis,
	}

	// Call Service
	results, totalData, errServ := cont.Service.GetAllByPengusulId(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	data, errData := pagination.ResponseToJson(pagination.Response{
		Data:      results,
		TotalData: totalData,
	})
	if errData != nil {
		exception.ErrorHandler(context, errData)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UsulanControllerImpl) Delete(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	errServ := cont.Service.Delete(idFinal)
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
