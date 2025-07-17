package controller

import "github.com/gin-gonic/gin"

type UsulanController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetByJenisAndStatus(context *gin.Context)
	SearchByJudul(context *gin.Context)
	Filter(context *gin.Context)
	UpdateStatus(context *gin.Context)
	GetAll(context *gin.Context)
	GetAllValidated(context *gin.Context)
	GetAllByPengusulId(context *gin.Context)
	Delete(context *gin.Context)
}
