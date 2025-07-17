package controller

import "github.com/gin-gonic/gin"

type FileController interface {
	UploadFile(context *gin.Context)
	DeleteFile(context *gin.Context)
}
