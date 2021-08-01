package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "Home Page",
		"login":  true,
		"logout": false,
	})
}
