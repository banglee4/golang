package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	RegisterAdmin(context *gin.Context)
	RegisterDosen(context *gin.Context)
	RegisterUppm(context *gin.Context)
	RegisterLppm(context *gin.Context)
	LoginAdmin(context *gin.Context)
	Login(context *gin.Context)
	ResetPassword(context *gin.Context)
	UpdatePassword(context *gin.Context)
	GetDosenByUserId(context *gin.Context)
	GetUppmByUserId(context *gin.Context)
	GetLppmByUserId(context *gin.Context)
	GetAllDosen(context *gin.Context)
	GetAllUppm(context *gin.Context)
	GetAllLppm(context *gin.Context)
	SearchDosenByName(context *gin.Context)
	SearchUppmByName(context *gin.Context)
	SearchLppmByName(context *gin.Context)
	Delete(context *gin.Context)
}
