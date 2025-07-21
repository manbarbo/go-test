package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterLifeCheckRoutes(router *gin.Engine) {
	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})
}
