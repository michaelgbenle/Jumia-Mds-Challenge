package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"math"
)

func GetProductSku(sku string) models.Product {
	product := models.Product{}
	Db.Where("sku= ?", sku).First(product)
	return product
}

func consumeStock(product models.Product) models.Order {
	initialStock := int(math.Abs(float64(product.Stock)))
	Db.First()

}
