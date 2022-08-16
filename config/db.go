package config

import (
	"gdevcha/main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/tutorial"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Usuario{})
	DB = db
}
