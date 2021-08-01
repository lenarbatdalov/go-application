package main

import (
	"github.com/lenarbatdalov/go-application/app"
)

func main() {
	container := app.BuildContainer()
	err := container.Invoke(func(s *app.Server) {
		s.Run()
	})

	if err != nil {
		panic(err)
	}
}
