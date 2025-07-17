package controller

import (
	"net/http"
	"sippm/exception"
	"sippm/helper"
	"sippm/model/service"
	"sippm/model/web"

	"github.com/gin-gonic/gin"
)

type FileControllerImpl struct {
	Serv service.FileService
}

func NewFileControllerImpl(serv service.FileService) *FileControllerImpl {
	return &FileControllerImpl{Serv: serv}
}

func (cont *FileControllerImpl) UploadFile(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		exception.ErrorHandler(context, err)
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		exception.ErrorHandler(context, "no files uploaded")
		return
	}

	uploadedFiles, errServ := cont.Serv.UploadFile(files)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   uploadedFiles,
	}

	if err := helper.WriteToResponseBody(context, api.Code, api); err != nil {
		exception.ErrorHandler(context, err)
		return
	}
}

func (cont *FileControllerImpl) DeleteFile(context *gin.Context) {
	url := context.Query("url")

	// Call Service
	err := cont.Serv.DeleteFile(url)
	if err != nil {
		exception.ErrorHandler(context, err)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	if err := helper.WriteToResponseBody(context, api.Code, api); err != nil {
		exception.ErrorHandler(context, err)
		return
	}
}
