package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	Password     string
	LoginDetails LoginDetails `gorm:"foreignKey:Email"`
}
