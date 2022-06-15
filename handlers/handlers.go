package handlers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/database"
	"github.com/michaelgbenle/jumiaMds/models"
	"log"
	"net/http"
)

type handler struct {
	DB database.DB
}

func HandleConstruct() *handler {
	return &handler{}
}

func (h handler) GetProductBySku(c *gin.Context) {
	sku := c.Query("sku")
	country := c.Query("country")
	product := h.DB.GetProductSku(sku, country)
	c.JSON(http.StatusOK, gin.H{
		"message": product,
	})
}

func (h *handler) ConsumeStock(c *gin.Context) {
	product := models.Product{}
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding json",
		})
	}
	stockSold := h.DB.SellStock(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": stockSold,
	})
}

func (h *handler) BulkUploadFromCsv(c *gin.Context) {
	csvFile, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file upload error",
		})
	}

	buf, err := csvFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file error",
		})
	}
	defer buf.Close() //needs review

	reader := csv.NewReader(buf)

	//reader.Comma = ','

	reader.LazyQuotes = true

	csvLines, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}

	h.DB.BulkUpload(csvLines)

	c.JSON(http.StatusOK, gin.H{
		"message": "Bulk update successful",
	})
}
