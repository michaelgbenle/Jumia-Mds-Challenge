package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Country string `json:"country"`
	Sku     string `gorm:"index:idx_skucountry,priority:1" json:"sku"`
	Name    string `json:"name"`
	Stock   int    `json:"stock"`
}
