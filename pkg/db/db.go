package db

import (
	config "github.com/Suranjan77/go-manage-event/pkg/config"

	"github.com/Suranjan77/go-manage-event/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func SetupDB() {

	db, err := gorm.Open(mysql.Open(config.DBConnectionUrl()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	conn, _ := db.DB()
	conn.SetMaxIdleConns(config.MaxIdleConnections())
	conn.SetConnMaxLifetime(config.MaxConnectionLifeTimeDuration())
	conn.SetMaxOpenConns(config.MaxOpenConnections())

	automigrate(db)

	database = db
}

func GetDB() *gorm.DB {
	return database
}

func automigrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.LoginDetails{})
}
