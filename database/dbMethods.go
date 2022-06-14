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

	Db.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, initialStock)

	order := models.Order{
		ProductId: product.ID,
		Amount:    uint(initialStock),
	}

	if product.ID <= 0 {

	}

	return order
}
