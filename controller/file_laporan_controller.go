package controller

import "github.com/gin-gonic/gin"

type FileLaporanController interface {
	UpdateFileUsulan(context *gin.Context)
	UpdateStatusUsulan(context *gin.Context)
	UpdateCatatanUsulan(context *gin.Context)
	UpdateFileLaporanKemajuan(context *gin.Context)
	UpdateStatusLaporanKemajuan(context *gin.Context)
	UpdateCatatanLaporanKemajuan(context *gin.Context)
	UpdateFileLaporanAkhir(context *gin.Context)
	UpdateStatusLaporanAkhir(context *gin.Context)
	UpdateCatatanLaporanAkhir(context *gin.Context)
}
