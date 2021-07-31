package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/go-application/controller"
	"github.com/lenarbatdalov/go-application/database"
	"github.com/lenarbatdalov/go-application/repository"
	"github.com/lenarbatdalov/go-application/service"
)

var (
	sessionSecretKey = "secret"
	store            = cookie.NewStore([]byte(sessionSecretKey))

	db                                  = database.InitConnect()
	userRepository                      = repository.NewUserRepository(db)
	loginService   service.LoginService = service.NewLoginService(userRepository)

	loginController controller.LoginController = controller.NewLoginController(loginService)
)

func main() {
	server := gin.New()

	server.Use(sessions.Sessions("mysession", store))
	server.LoadHTMLGlob("templates/**/*.html")
	server.Static("/assets", "./assets")

	server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title":  "Home Page",
			"login":  true,
			"logout": false,
		})
	})

	server.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login", gin.H{"title": "Login Page"})
	})
	server.POST("/login", loginController.Login)

	err := server.Run(":8080")
	if err != nil {
		panic("server not started")
	}
}
