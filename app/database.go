package app

import (
	"sippm/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection(username string, password string, addr string, name string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + addr + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialect := mysql.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, dbErr := db.DB()
	helper.FatalError(dbErr)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
