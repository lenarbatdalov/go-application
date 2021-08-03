package entity

import "gorm.io/gorm"

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(255)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(255)"`
}

func CreateUserSeed(db *gorm.DB) {
	var user User
	db.Where("username = ? AND password = ?", "admin", "admin").First(&user)
	if user.ID == 0 {
		db.Create(&User{Username: "admin", Password: "admin"})
	}
}
