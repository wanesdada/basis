package httpserver

import (
	"github.com/gin-gonic/gin"
	"qp_web_server/config"
	"qp_web_server/httpserver/servermiddleware"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.GetService().Mode)
	gin.ForceConsoleColor()
	router := gin.Default()
	//各种中间件
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(servermiddleware.BaseLogger())
	router.Use(servermiddleware.EnableCors([]string{"*"}))
	initRoutes(router)
	return router
}
