package entity

import "time"

type Key struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserId      int
	User        User      `gorm:"foreignKey:UserId"`
	Url         string    `json:"url" binding:"required" gorm:"type:varchar(255)"`
	Login       string    `json:"login" binding:"required" gorm:"type:varchar(255)"`
	Password    string    `json:"password" binding:"required" gorm:"type:varchar(255)"`
	Description string    `json:"description" binding:"required" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
