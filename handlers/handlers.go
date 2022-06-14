package handlers

import "github.com/gin-gonic/gin"

func SampleRequest(c *gin.Context) {

	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func GetProductBySku(c *gin.Context){
	sku:=c.Query("sku")
	countrCode :=c.Query("country")
	products:=
}