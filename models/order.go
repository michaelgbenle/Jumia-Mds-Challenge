package models

import (
	"gorm.io/gorm"
)

type order struct {
	gorm.Model
	ProductId uint
	Amount    uint
	Product   Product
}
