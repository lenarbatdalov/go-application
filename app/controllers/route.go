package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func generateToken() string {
	// сгенерировать токен
	// положить токен в файл конфига
	// вернуть токен для подписи сессии
	return "secret"
}

func Server() *gin.Engine {
	sessionSecretKey := generateToken()

	server := gin.New()

	store := cookie.NewStore([]byte(sessionSecretKey))
	server.Use(sessions.Sessions("mysession", store))

	server.Static("/css", "/templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	return server
}
