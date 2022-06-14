package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/database"
	"github.com/michaelgbenle/jumiaMds/models"
	"net/http"
)

func SampleRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func GetProductBySku(c *gin.Context) {
	sku := c.Query("sku")
	product := database.GetProductSku(sku)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func ConsumeStock(c *gin.Context) {
	product := &models.Product{}
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding json",
		})
	}
	//result := services.MakeOrder(product)
	//return ctx.Status(http.StatusOK).JSON(result)
}
