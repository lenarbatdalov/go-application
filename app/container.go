package app

import (
	"github.com/lenarbatdalov/go-application/controller"
	"github.com/lenarbatdalov/go-application/repository"
	"github.com/lenarbatdalov/go-application/service"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(InitConnect)
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewLoginService)
	container.Provide(controller.NewLoginController)
	container.Provide(NewServer)

	return container
}
