package entity

import "time"

/** Приход */
type Income struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserId      int       `json:"user_id"`
	User        User      `gorm:"foreignKey:UserId"`
	CategoryId  int       `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryId"`
	AccountId   int       `json:"account_id"`
	Account     Account   `gorm:"foreignKey:AccountId"`
	Amount      string    `json:"amount" binding:"required" gorm:"type:decimal(10,2)"`
	Description string    `json:"description" binding:"required" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
