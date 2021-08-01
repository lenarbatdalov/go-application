package service

import "github.com/lenarbatdalov/go-application/repository"

type LoginService interface {
	FindUser(username string, password string) bool
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(ur repository.UserRepository) LoginService {
	return &loginService{
		userRepository: ur,
	}
}

func (ls *loginService) FindUser(username string, password string) bool {
	return true
}
