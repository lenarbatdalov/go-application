package database

import (
	"github.com/lenarbatdalov/go-application/cmd"
	"github.com/lenarbatdalov/go-application/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitConnect() *gorm.DB {
	runCommand()

	db, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migration(db)
	seeding(db)

	return db
}

func runCommand() {
	cmd.DefineDatabase()
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

func seeding(db *gorm.DB) {
	db.Create(&entity.User{Username: "admin", Password: "admin"})
}
