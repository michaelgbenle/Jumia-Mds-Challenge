package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Sku     string `gorm:"index:idx_member,priority:1" json:"sku"`
	Country string `gorm:"index:idx_member,priority:2" json:"country"`
	Name    string `json:"name"`
	Stock   int    `json:"stock"`
}
