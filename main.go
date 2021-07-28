package main

import (
	"github.com/lenarbatdalov/go-application/service"
)

var (
	// db = database.InitConnect()
	s = service.Server()
)

func main() {
	err := s.Run(":8080")
	if err != nil {
		panic("server not started")
	}
}
