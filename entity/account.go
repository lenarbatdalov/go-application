package entity

import "time"

/** Счет */
type Account struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Type      string    `json:"amount" binding:"required" gorm:"type:varchar(255)"`
	Icon      string    `json:"icon" binding:"required" gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
