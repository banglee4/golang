package app

import (
	"net/http"
	"sippm/controller"
	"sippm/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserCont         controller.UserController
	DosenCont        controller.DosenController
	UppmCont         controller.UppmController
	LppmCon          controller.LppmController
	UsulanCont       controller.UsulanController
	FileLaporanCont  controller.FileLaporanController
	FileCont         controller.FileController
	PengumumanCont   controller.PengumumanController
	JabatanKetuaCont controller.JabatanController

	SecretKey string
	Engine    *gin.Engine
}

func NewRouter(r Router) *Router {
	// CORS middleware
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Route
	userGroup := r.Engine.Group("/user")
	userGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		userGroup.POST("/admin/register", r.UserCont.RegisterAdmin)
		userGroup.POST("/admin/login", r.UserCont.LoginAdmin)

		userGroup.POST("/dosen/register", r.UserCont.RegisterDosen)
		userGroup.GET("/dosen/:userId", r.UserCont.GetDosenByUserId)
		userGroup.GET("/dosen", r.UserCont.GetAllDosen)
		userGroup.PUT("/dosen", r.DosenCont.UpdateData)

		userGroup.POST("/uppm/register", r.UserCont.RegisterUppm)
		userGroup.GET("/uppm/:userId", r.UserCont.GetUppmByUserId)
		userGroup.GET("/uppm", r.UserCont.GetAllUppm)
		userGroup.PUT("/uppm", r.UppmCont.UpdateData)

		userGroup.POST("/lppm/register", r.UserCont.RegisterLppm)
		userGroup.GET("/lppm/:userId", r.UserCont.GetLppmByUserId)
		userGroup.GET("/lppm", r.UserCont.GetAllLppm)
		userGroup.PUT("/lppm", r.LppmCon.UpdateData)

		userGroup.POST("/login", r.UserCont.Login)
		userGroup.PUT("/password", r.UserCont.UpdatePassword)
		userGroup.PUT("/reset-password", r.UserCont.ResetPassword)
		userGroup.DELETE("/:userId", r.UserCont.Delete)
	}

	usulanGroup := r.Engine.Group("/usulan")
	usulanGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		usulanGroup.POST("/", r.UsulanCont.Create)
		usulanGroup.PUT("/", r.UsulanCont.Update)
		usulanGroup.PUT("/status", r.UsulanCont.UpdateStatus)
		usulanGroup.GET("/dosen", r.UsulanCont.GetAllByPengusulId)
		usulanGroup.GET("", r.UsulanCont.GetAll)
		usulanGroup.GET("/selesai", r.UsulanCont.GetAllValidated)
		usulanGroup.DELETE("/:id", r.UsulanCont.Delete)
	}

	fileLaporanGroup := r.Engine.Group("/file-laporan")
	fileLaporanGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		fileLaporanGroup.PUT("/file-usulan", r.FileLaporanCont.UpdateFileUsulan)
		fileLaporanGroup.PUT("/catatan-usulan", r.FileLaporanCont.UpdateCatatanUsulan)
		fileLaporanGroup.PUT("/status-usulan", r.FileLaporanCont.UpdateStatusUsulan)

		fileLaporanGroup.PUT("/file-laporan-kemajuan", r.FileLaporanCont.UpdateFileLaporanKemajuan)
		fileLaporanGroup.PUT("/catatan-laporan-kemajuan", r.FileLaporanCont.UpdateCatatanLaporanKemajuan)
		fileLaporanGroup.PUT("/status-laporan-kemajuan", r.FileLaporanCont.UpdateStatusLaporanKemajuan)

		fileLaporanGroup.PUT("/file-laporan-akhir", r.FileLaporanCont.UpdateFileLaporanAkhir)
		fileLaporanGroup.PUT("/catatan-laporan-akhir", r.FileLaporanCont.UpdateCatatanLaporanAkhir)
		fileLaporanGroup.PUT("/status-laporan-akhir", r.FileLaporanCont.UpdateStatusLaporanAkhir)
	}

	fileGroup := r.Engine.Group("/file")
	fileGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		fileGroup.POST("/upload", r.FileCont.UploadFile)
		fileGroup.DELETE("/delete", r.FileCont.DeleteFile)
	}

	pengGroup := r.Engine.Group("/pengumuman")
	pengGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		pengGroup.POST("/", r.PengumumanCont.Create)
		pengGroup.PUT("/", r.PengumumanCont.Update)
		pengGroup.GET("/", r.PengumumanCont.GetAll)
		pengGroup.DELETE("/:id", r.PengumumanCont.Delete)
	}

	jabatanGroup := r.Engine.Group("/jabatan-ketua")
	jabatanGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		jabatanGroup.POST("/", r.JabatanKetuaCont.Create)
		jabatanGroup.PUT("/", r.JabatanKetuaCont.Update)
		jabatanGroup.GET("/", r.JabatanKetuaCont.GetAll)
		jabatanGroup.DELETE("/:id", r.JabatanKetuaCont.Delete)
	}

	r.Engine.StaticFS("/assets", http.Dir("./assets"))

	return &Router{
		UserCont:         r.UserCont,
		DosenCont:        r.DosenCont,
		UppmCont:         r.UppmCont,
		LppmCon:          r.LppmCon,
		UsulanCont:       r.UsulanCont,
		FileLaporanCont:  r.FileLaporanCont,
		FileCont:         r.FileCont,
		PengumumanCont:   r.PengumumanCont,
		JabatanKetuaCont: r.JabatanKetuaCont,

		SecretKey: r.SecretKey,
		Engine:    r.Engine,
	}
}
