package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Country string
	Sku     string
	Name    string
	Stock   uint
}
