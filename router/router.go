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
		log.Println(err)
	}

	apirouter := router.Group("/api/v1")
	apirouter.GET("/product", handlers.GetProductBySku)
	apirouter.Post("/product/consume", handlers.ConsumeStock)
	apirouter.POST("/product/bulkupdate", handlers.BulkUploadFromCsv)

	return router
}
