package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetProductSku(sku string) models.Product {
	product := models.Product{}
	Db.Where()
	return product
}
