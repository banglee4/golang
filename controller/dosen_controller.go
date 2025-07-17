package controller

import "github.com/gin-gonic/gin"

type DosenController interface {
	UpdateData(context *gin.Context)
}
