package httpserver

import (
	"basis/httpserver/handle"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/version", handle.Version)
	router.GET("/ping", handle.Ping)
}
