package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(port string, db *gorm.DB) {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.GET("/")
	router.POST("")
	router.POST("")
	router.POST("")
	router.POST("")
	router.POST("")
}
