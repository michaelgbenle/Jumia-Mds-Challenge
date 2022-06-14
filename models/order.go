package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductId uint
	Amount    uint
	Product   Product
}
