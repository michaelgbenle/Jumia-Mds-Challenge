package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/database"
)

func SampleRequest(c *gin.Context) {

	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func GetProductBySku(c *gin.Context) {
	sku := c.Query("sku")
	products := database.GetProductSku(sku)
}
