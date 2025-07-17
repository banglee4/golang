package main

import (
	"os"
	"sippm/app"
	"sippm/controller"
	"sippm/helper"
	"sippm/model/repository"
	"sippm/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Load .env file
	err := godotenv.Load(".env")
	helper.FatalError(err)

	// Other Env
	secretKey := os.Getenv("SECRET_KEY")
	port := os.Getenv("APP_PORT")

	// Database config
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Open DB Connection
	db := app.OpenConnection(dbUser, dbPass, dbHost+":"+dbPort, dbName)

	// Initialize Repo
	userRepo := repository.NewUserRepositoryImpl(db)
	dosenRepo := repository.NewDosenRepositoryImpl(db)
	uppmRepo := repository.NewUppmRepositoryImpl(db)
	lppmRepo := repository.NewLppmRepositoryImpl(db)
	usulanRepo := repository.NewUsulanRepositoryImpl(db)
	fileLaporanRepo := repository.NewFileLaporanRepositoryImpl(db)
	pengRepo := repository.NewPengumumanRepositoryImpl()
	pengDokRepo := repository.NewPengDokumenRepositoryImpl()
	jabatanKetuaRepo := repository.NewJabatanRepositoryImpl()

	// Initialize Service
	userServ := service.NewUserServiceImpl(userRepo, secretKey)
	dosenServ := service.NewDosenServiceImpl(dosenRepo)
	uppmServ := service.NewUppmServiceImpl(uppmRepo)
	lppmServ := service.NewLppmServiceImpl(lppmRepo)
	usulanServ := service.NewUsulanServiceImpl(usulanRepo)
	fileLaporanServ := service.NewFileLaporanServiceImpl(fileLaporanRepo)
	fileServ := service.NewFileServImpl()
	pengServ := service.NewPengumumanServiceImpl(pengRepo, pengDokRepo, db)
	jabatanServ := service.NewJabatanKetuaServiceImpl(jabatanKetuaRepo, db)

	// Initialize Controller
	userCont := controller.NewUserControllerImpl(userServ)
	dosenCont := controller.NewDosenControllerImpl(dosenServ)
	uppmCont := controller.NewUppmControllerImpl(uppmServ)
	lppmCont := controller.NewLppmControllerImpl(lppmServ)
	usulanCont := controller.NewUsulanControllerImpl(usulanServ)
	fileLaporanCont := controller.NewFileLaporanControllerImpl(fileLaporanServ)
	fileCont := controller.NewFileControllerImpl(fileServ)
	pengCont := controller.NewPengumumanControllerImpl(pengServ)
	jabatanKetuaCont := controller.NewJabatanControllerImpl(jabatanServ)

	// Setup router
	engine := gin.Default()
	routerParent := app.Router{
		UserCont:         userCont,
		DosenCont:        dosenCont,
		UppmCont:         uppmCont,
		LppmCon:          lppmCont,
		UsulanCont:       usulanCont,
		FileLaporanCont:  fileLaporanCont,
		FileCont:         fileCont,
		PengumumanCont:   pengCont,
		JabatanKetuaCont: jabatanKetuaCont,

		SecretKey: secretKey,
		Engine:    engine,
	}
	router := app.NewRouter(routerParent)

	// Run Server
	if port == "" {
		port = ":3001" // Default fallback
	}
	err = router.Engine.Run(port)
	helper.FatalError(err)
}
