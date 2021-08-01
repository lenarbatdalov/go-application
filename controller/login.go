package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/go-application/service"
)

type LoginController interface {
	LoginPage(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
}

func NewLoginController(ls service.LoginService) LoginController {
	return &loginController{
		loginService: ls,
	}
}

func (lс *loginController) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{"title": "Login Page"})
}

func (lс *loginController) Login(ctx *gin.Context) {
	username := ctx.PostForm("login")
	password := ctx.PostForm("password")
	// remember := ctx.PostForm("remember")

	check := lс.loginService.FindUser(username, password)
	if check {
		ctx.Redirect(http.StatusFound, "/")
	} else {
		ctx.HTML(http.StatusOK, "login", gin.H{"title": "Login Page", "err": "Не верно введен логин или пароль."})
	}
}
