package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DefaultController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userLogin := session.Get("user") == true
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "Home Page",
		"login":  !userLogin,
		"logout": userLogin,
	})
}
