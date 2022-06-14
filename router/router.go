package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/")
	apirouter.POST("")
	apirouter.POST("")
	apirouter.POST("")
	apirouter.POST("")
	apirouter.POST("")

	return router
}
