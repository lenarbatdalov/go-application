package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", nil)
}
