package database

import (
	"github.com/lenarbatdalov/go-application/app/commands"
	"github.com/lenarbatdalov/go-application/database/entities"
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
	commands.DefineDatabase()
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}

func seeding(db *gorm.DB) {
	db.Create(&entities.User{Username: "admin", Password: "admin"})
}
