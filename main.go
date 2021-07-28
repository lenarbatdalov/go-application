package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	sessionSecretKey = "secret"
	store            = cookie.NewStore([]byte(sessionSecretKey))
	// db = database.InitConnect()
)

func main() {
	server := gin.New()

	server.Use(sessions.Sessions("mysession", store))
	server.LoadHTMLGlob("templates/**/*.html")
	server.Static("/assets", "./assets")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Home Page",
			"name":  "Lenar",
		})
	})

	err := server.Run(":8080")
	if err != nil {
		panic("server not started")
	}
}
