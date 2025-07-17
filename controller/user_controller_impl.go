package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"
	"sippm/model/web/pagination"
	"sippm/model/web/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserControllerImpl(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{Service: service}
}

func (cont *UserControllerImpl) RegisterAdmin(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterAdminRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.RegisterAdmin(request)
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

func (cont *UserControllerImpl) RegisterDosen(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterDosenRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.RegisterDosen(request)
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

func (cont *UserControllerImpl) RegisterUppm(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterUppmRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.RegisterUppm(request)
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

func (cont *UserControllerImpl) RegisterLppm(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterLppmRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.RegisterLppm(request)
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

func (cont *UserControllerImpl) LoginAdmin(context *gin.Context) {
	// Parse Request Body
	request := user.LoginRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	token, errServ := cont.Service.LoginAdmin(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   token,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserControllerImpl) Login(context *gin.Context) {
	// Parse Request Body
	request := user.LoginRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	token, errServ := cont.Service.Login(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   token,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserControllerImpl) ResetPassword(context *gin.Context) {
	// Parse Request Body
	request := user.ResetPwRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.ResetPassword(request)
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

func (cont *UserControllerImpl) UpdatePassword(context *gin.Context) {
	// Parse Request Body
	request := user.UpdatePwRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Service.UpdatePassword(request)
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

func (cont *UserControllerImpl) GetDosenByUserId(context *gin.Context) {
	id := context.Param("userId")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	result, errServ := cont.Service.GetDosenByUserId(idFinal)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserControllerImpl) GetUppmByUserId(context *gin.Context) {
	id := context.Param("userId")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	result, errServ := cont.Service.GetUppmByUserId(idFinal)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserControllerImpl) GetLppmByUserId(context *gin.Context) {
	id := context.Param("userId")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	result, errServ := cont.Service.GetLppmByUserId(idFinal)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserControllerImpl) GetAllDosen(context *gin.Context) {
	// Call Service
	results, totalData, errServ := cont.Service.GetAllDosen()
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

func (cont *UserControllerImpl) GetAllUppm(context *gin.Context) {
	// Call Service
	results, totalData, errServ := cont.Service.GetAllUppm()
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

func (cont *UserControllerImpl) GetAllLppm(context *gin.Context) {
	// Call Service
	results, totalData, errServ := cont.Service.GetAllLppm()
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

func (cont *UserControllerImpl) SearchDosenByName(context *gin.Context) {
	// Parse Request Body
	name := context.Query("nama")

	// Call Service
	results, errServ := cont.Service.SearchDosenByName(name)
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

func (cont *UserControllerImpl) SearchUppmByName(context *gin.Context) {
	// Parse Request Body
	name := context.Query("nama")

	// Call Service
	results, errServ := cont.Service.SearchUppmByName(name)
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

func (cont *UserControllerImpl) SearchLppmByName(context *gin.Context) {
	// Parse Request Body
	name := context.Query("nama")

	// Call Service
	results, errServ := cont.Service.SearchLppmByName(name)
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

func (cont *UserControllerImpl) Delete(context *gin.Context) {
	id := context.Param("userId")
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
