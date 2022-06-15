package database

import "github.com/michaelgbenle/jumiaMds/models"

type DB interface {
	GetProductSku(sku, country string) (*models.Product, error)
	SellStock(product *models.Product) models.Order
	SellOrCreate(product *models.Product)
	ProductCreate(product *models.Product) *models.Product
	BulkUpload(file [][]string)
}
