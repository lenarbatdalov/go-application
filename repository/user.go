package repository

import (
	"github.com/lenarbatdalov/go-application/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Save(user entity.User)
	// Update(user entity.User)
	// Delete(user entity.User)
	FindUser(name string) entity.User
	// FindAll() []entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (ur *userRepository) FindUser(name string) entity.User {
	var user entity.User
	ur.connection.Set("gorm:auto_preload", true).First(&user, "username = ?", name)
	// fmt.Println(user)
	return user
}
