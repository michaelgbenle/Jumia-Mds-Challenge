package router

import (
	"github.com/gin-gonic/gin"
	"os"
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

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router
}
