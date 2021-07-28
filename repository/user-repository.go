package repository

import "gorm.io/gorm"

type UserRepository interface {
	// Save(user entity.User)
	// Update(user entity.User)
	// Delete(user entity.User)
	// Find(name string) entity.User
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
