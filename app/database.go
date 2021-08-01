package app

import (
	"os"

	"github.com/lenarbatdalov/go-application/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func defineDatabase() {
	fileName := "database/database.sqlite"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

func seeding(db *gorm.DB) {
	db.Create(&entity.User{Username: "admin", Password: "admin"})
}

func InitConnect() *gorm.DB {
	defineDatabase()

	db, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migration(db)
	seeding(db)

	return db
}
