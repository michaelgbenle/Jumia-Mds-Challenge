package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func GetProductSku(sku, country string) models.Product {
	product := models.Product{}
	Db.Where("sku= ? AND country=?", sku, country).First(product)
	return product
}

func SellStock(product *models.Product) models.Order {
	purchaseStock := int(math.Abs(float64(product.Stock)))

	Db.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, purchaseStock)

	order := models.Order{
		ProductId: product.ID,
		Quantity:  uint(purchaseStock),
	}

	if product.ID <= 0 {
		log.Println("Product Not Available")
		return order
	}

	//create the order
	Db.Create(&order)

	//Update product amount to reflect change
	Db.Model(models.Product{}).Where("id = ?", product.ID).Updates(models.Product{Stock: product.Stock - purchaseStock})

	return order
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
		trans.Rollback()
		return product
	}
	product.Stock += changeInStock
	trans.Commit()
	return product
}

func BulkUpload(file [][]string) {
	dbChan := make(chan int, 90)
	for i, fileLine := range file {
		if i == 0 {
			continue
		}
		result := strings.ReplaceAll(fileLine[0], `","`, "!")
		resultArray := strings.Split(result, "!")
		if len(resultArray) != 4 {
			log.Println("incomplete data")
			continue
		}
		stock, err := strconv.Atoi(resultArray[3])
		if err != nil {
			log.Println("incorrect value")
			continue
		}
		product := &models.Product{
			Country: resultArray[0],
			Sku:     resultArray[1],
			Name:    resultArray[2],
			Stock:   stock,
		}
		wg.Add(1)
		dbChan <- 1
		go func(product models.Product) {
			SwitchSellBuy(&product)
			wg.Done()
			<-dbChan

		}(*product)
	}
	//wg.Wait()
}
