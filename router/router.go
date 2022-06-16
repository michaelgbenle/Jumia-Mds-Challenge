package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/handlers"
)

func SetupRouter() *gin.Engine {
	handler := handlers.HandleConstruct()
	router := gin.Default()
	//PDB := new(database.PostgresDb)
	//err := PDB.SetupDb()
	//if err != nil {
	//	log.Println(err)
	//}

	apirouter := router.Group("/api/v1")
	apirouter.GET("/product", handler.GetProductBySku)
	apirouter.POST("/product/consume", handler.ConsumeStock)
	apirouter.POST("/product/bulkupdate", handler.BulkUploadFromCsv)

	return router
}
