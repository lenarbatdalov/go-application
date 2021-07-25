package main

import (
	"github.com/lenarbatdalov/go-application/app/controllers"
)

var (
	// db = database.InitConnect()
	s = controllers.Server()
)

func main() {
	s.Run(":8080")
}
