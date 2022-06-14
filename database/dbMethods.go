package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
)

func GetProductSku(sku string) models.Product {
	product := models.Product{}
	Db.Where("sku= ?", sku)
	return product
}
