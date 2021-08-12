package app

import (
	"os"

	"github.com/lenarbatdalov/go-application/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func defineDatabase() {
	fileName := "database.sqlite"
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
	db.AutoMigrate(&entity.Key{})
	db.AutoMigrate(&entity.Account{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Expense{})
	db.AutoMigrate(&entity.Income{})
}

func seeding(db *gorm.DB) {
	entity.CreateUserSeed(db)
	// db.Create(&entity.Key{Login: "admin", Password: "admin"})
}

func InitConnect() *gorm.DB {
	defineDatabase()

	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migration(db)
	seeding(db)

	return db
}
