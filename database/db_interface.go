package database

import "github.com/michaelgbenle/jumiaMds/models"

//DB interface implements database methods
type DB interface {
	GetProductSku(sku, country string) (*models.Product, error)
	SellStock(product *models.Product) (*models.Order, error)
	SellOrCreate(product *models.Product)
	ProductCreate(product *models.Product) *models.Product
	BulkUpload(file [][]string)
}
