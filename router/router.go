package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/handlers"
)

func SetupRouter(handler handlers.Handler) *gin.Engine {
	router := gin.Default()

	apirouter := router.Group("/api/v1")
	apirouter.GET("/product", handler.GetProductBySku)
	apirouter.POST("/product/consume", handler.ConsumeStock)
	apirouter.POST("/product/bulkupdate", handler.BulkUploadFromCsv)

	return router
}
