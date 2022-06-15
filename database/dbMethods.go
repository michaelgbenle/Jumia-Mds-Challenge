package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"log"
	"math"
	"sync"
)

var wg sync.WaitGroup

func GetProductSku(sku, country string) models.Product {
	product := models.Product{}
	Db.Where("sku= ? AND country=?", sku, country).First(product)
	return product
}

func SellStock(product *models.Product) models.Order {
	initialStock := int(math.Abs(float64(product.Stock)))

	Db.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, initialStock)

	order := models.Order{
		ProductId: product.ID,
		Amount:    uint(initialStock),
	}

	if product.ID <= 0 {
		log.Println("Product Not Available")
		return order
	}

	//create the order
	Db.Create(&order)

	//Update product amount to reflect change
	Db.Model(models.Product{}).Where("id = ?", product.ID).Updates(models.Product{Stock: product.Stock - initialStock})

	return order
}

func BulkUpload(products *[]models.Product) {
	dbconnections := make(chan int, 90)
	for _, product := range *products {
		wg.Add(1)
		dbconnections <- 1
		go func(product models.Product) {
			SwitchSellBuy(&product)
			wg.Done()
			<-dbconnections
		}(product)
	}
}

func SwitchSellBuy(product *models.Product) {
	if int(product.Stock) < 0 {
		SellStock(product)
	} else {
		ProductCreate(product)
	}
}
func ProductCreate(product *models.Product) *models.Product {
	changeInStock := product.Stock
	trans := Db.Begin()
	trans.Where("sku =? AND country = ?", product.Sku, product.Country).First(product)

	if product.ID == 0 {
		if err := trans.Create(product).Error; err != nil {
			trans.Rollback()
			return product
		}
		trans.Commit()
		return product
	}
	//update product to reflect change
	if err := trans.Model(models.Product{}).
		Where("id = ?", product.ID).Updates(models.Product{Stock: product.Stock + changeInStock}).Error; err != nil {

	}
}
