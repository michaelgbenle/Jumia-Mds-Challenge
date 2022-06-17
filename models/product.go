package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model
	Country string `gorm:"index:idx_member,priority:2" json:"country"`
	Sku     string `gorm:"index:idx_member,priority:1" json:"sku"`
	Name    string `json:"name"`
	Stock   int    `json:"stock"`
}
