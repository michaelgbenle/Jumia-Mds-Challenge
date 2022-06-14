package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Country string `json:"country"`
	Sku     string `json:"sku"`
	Name    string `json:"name"`
	Stock   int    `json:"stock"`
}
