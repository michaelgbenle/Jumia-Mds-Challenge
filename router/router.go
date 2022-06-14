package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/")

	return router
}
