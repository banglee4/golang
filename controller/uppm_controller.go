package controller

import "github.com/gin-gonic/gin"

type UppmController interface {
	UpdateData(context *gin.Context)
}
