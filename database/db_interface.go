package database

import "github.com/michaelgbenle/jumiaMds/models"

type DB interface {
	GetProductSku(sku, country string) models.Product
}
