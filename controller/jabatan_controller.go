package controller

import (
	"github.com/gin-gonic/gin"
)

type JabatanController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetAll(context *gin.Context)
	Delete(context *gin.Context)
}
