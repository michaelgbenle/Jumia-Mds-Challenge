package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/database"
	"github.com/michaelgbenle/jumiaMds/handlers"
	"log"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	err := database.SetupDb()
	if err != nil {
		log.Fatalln(err)
	}

	apirouter := router.Group("/api/v1")

	apirouter.GET("/", handlers.SampleRequest)
	apirouter.GET("/product", handlers.GetProductBySku)
	apirouter.GET("/product/consume", handlers.ConsumeStock)

	return router
}
