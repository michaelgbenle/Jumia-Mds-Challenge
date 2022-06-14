package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Country string `json:"name"`
	Sku     string `json:"sku"`
	Name    string `json:"name"`
	Stock   uint   `json:"stock"`
}
