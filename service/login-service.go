package service

import "github.com/lenarbatdalov/go-application/repository"

type LoginService interface {
	Login(login string, password string) bool
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(userRepository repository.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (service *loginService) Login(login string, password string) bool {
	user := service.userRepository.Find(login)
	return user.Username == login
}
