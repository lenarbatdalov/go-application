package entities

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(255)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(255)"`
}
