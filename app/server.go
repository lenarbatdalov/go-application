package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/go-application/controller"
)

type Server struct {
	loginController controller.LoginController
}

func NewServer(loginController controller.LoginController) *Server {
	return &Server{
		loginController: loginController,
	}
}

func (s *Server) Run() {
	r := gin.New()

	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/assets", "./assets")

	auth := r.Group("/")
	{
		auth.GET("/", controller.DefaultController)
		auth.GET("/login", s.loginController.LoginPage)
		auth.POST("/login", s.loginController.Login)
		// auth.GET("/logout")
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
