package controller

import "github.com/gin-gonic/gin"

type LppmController interface {
	UpdateData(context *gin.Context)
}
