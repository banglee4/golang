package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/file_laporan"

	"github.com/gin-gonic/gin"
)

type FileLaporanControllerImpl struct {
	Service service.FileLaporanService
}

func NewFileLaporanControllerImpl(service service.FileLaporanService) *FileLaporanControllerImpl {
	return &FileLaporanControllerImpl{Service: service}
}

func (cont *FileLaporanControllerImpl) UpdateFileUsulan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateFileUsulanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateFileUsulan(request)
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

func (cont *FileLaporanControllerImpl) UpdateCatatanUsulan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateCatatanUsulanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateCatatanUsulan(request)
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

func (cont *FileLaporanControllerImpl) UpdateFileLaporanKemajuan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateFileLaporanKemajuanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateFileLaporanKemajuan(request)
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

func (cont *FileLaporanControllerImpl) UpdateCatatanLaporanKemajuan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateCatatanLaporanKemajuanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateCatatanLaporanKemajuan(request)
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

func (cont *FileLaporanControllerImpl) UpdateFileLaporanAkhir(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateFileLaporanAkhirRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateFileLaporanAkhir(request)
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

func (cont *FileLaporanControllerImpl) UpdateCatatanLaporanAkhir(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateCatatanLaporanAkhirRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateCatatanLaporanAkhir(request)
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

func (cont *FileLaporanControllerImpl) UpdateStatusUsulan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateStatusUsulanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateStatusUsulan(request)
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

func (cont *FileLaporanControllerImpl) UpdateStatusLaporanKemajuan(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateStatusLaporanKemajuanRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateStatusLaporanKemajuan(request)
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

func (cont *FileLaporanControllerImpl) UpdateStatusLaporanAkhir(context *gin.Context) {
	// Parse Request Body
	request := file_laporan.UpdateStatusLaporanAkhirRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdateStatusLaporanAkhir(request)
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
