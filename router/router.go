package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

func SetupRouter(port string, db *gorm.DB) {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

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
	return router, port
}
