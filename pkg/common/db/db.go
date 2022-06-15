package db

import (
	"github.com/Suranjan77/go-manage-event/pkg/common/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.LoginDetails{})

	return db
}
