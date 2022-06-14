package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"math"
)

func GetProductSku(sku, country string) models.Product {
	product := models.Product{}
	Db.Where("sku= ? AND country=?", sku, country).First(product)
	return product
}

func SellStock(product models.Product) models.Order {
	initialStock := int(math.Abs(float64(product.Stock)))
	Db.First(product, "sku=?", product.Sku)

	order := models.Order{
		ProductId: product.ID,
		Amount:    uint(initialStock),
	}

	return order
}
