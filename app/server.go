package app

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/assets", "./assets")

	auth := r.Group("/")
	{
		auth.GET("/", controller.DefaultController)
		auth.GET("/login", s.loginController.LoginPage)
		auth.POST("/login", s.loginController.Login)
		auth.GET("/logout", s.loginController.Logout)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
