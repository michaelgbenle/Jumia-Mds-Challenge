package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"log"
	"math"
	"sync"
)

var wg sync.WaitGroup

func GetProductSku(sku, country string) models.Product {
	product := models.Product{}
	Db.Where("sku= ? AND country=?", sku, country).First(product)
	return product
}

func SellStock(product *models.Product) models.Order {
	initialStock := int(math.Abs(float64(product.Stock)))

	Db.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, initialStock)

	order := models.Order{
		ProductId: product.ID,
		Amount:    uint(initialStock),
	}

	if product.ID <= 0 {
		log.Println("Product Not Available")
		return order
	}

	//create the order
	Db.Create(&order)

	//Update product amount to reflect change
	Db.Model(models.Product{}).Where("id = ?", product.ID).Updates(models.Product{Stock: product.Stock - initialStock})

	return order
}

func BulkUpload(products *[]models.Product) {
	dbconnections := make(chan int, 90)
	for _, product := range *products {
		wg.Add(1)
		dbconnections <- 1
	}
}
