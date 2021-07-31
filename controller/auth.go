package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/go-application/service"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
}

func NewLoginController(loginService service.LoginService) LoginController {
	return &loginController{
		loginService: loginService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) {
	login := ctx.PostForm("login")
	password := ctx.PostForm("password")
	// remember := ctx.PostForm("remember")

	if controller.loginService.Login(login, password) {
		ctx.Redirect(http.StatusFound, "/")
	} else {
		ctx.HTML(http.StatusOK, "login", gin.H{"title": "Login Page", "error": true})
	}
}
