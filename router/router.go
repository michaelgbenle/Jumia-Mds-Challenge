package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/", handlers.SampleRequest)

	return router
}
