package config

import (
	"Exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database,got err:%v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatalf("Failed to config database,got err:%v", err)
	}

	global.Db = db
}
