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
	country := c.Query("country")
	product := database.GetProductSku(sku, country)
	c.JSON(http.StatusOK, gin.H{
		"message": product,
	})
}

func ConsumeStock(c *gin.Context) {
	product := models.Product{}
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding json",
		})
	}
	stockSold := database.SellStock(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": stockSold,
	})
}

func BulkUpdate(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Bulk update successful",
	})

}
