package database

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

//PostgresDb implements the DB interface
type PostgresDb struct {
	DB *gorm.DB
}

func (pdb *PostgresDb) SetupDb() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pdb.DB = db

	err = pdb.DB.AutoMigrate(models.Product{}, models.Order{})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (pdb *PostgresDb) GetProductSku(sku, country string) models.Product {
	product := models.Product{}
	pdb.DB.Where("sku= ? AND country=?", sku, country).First(product)
	return product
}

func (pdb *PostgresDb) SellStock(product *models.Product) models.Order {
	purchaseStock := int(math.Abs(float64(product.Stock)))

	pdb.DB.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, purchaseStock)

	order := models.Order{
		ProductId: product.ID,
		Quantity:  uint(purchaseStock),
	}

	if product.ID <= 0 {
		log.Println("Product Not Available")
		return order
	}

	//create the order
	pdb.DB.Create(&order)

	//Update product amount to reflect change
	pdb.DB.Model(models.Product{}).Where("id = ?", product.ID).Updates(models.Product{Stock: product.Stock - purchaseStock})

	return order
}

func (pdb *PostgresDb) SellOrCreate(product *models.Product) {
	if int(product.Stock) < 0 {
		pdb.SellStock(product)
	} else {
		pdb.ProductCreate(product)
	}
}

func (pdb *PostgresDb) ProductCreate(product *models.Product) *models.Product {
	changeInStock := product.Stock
	trans := pdb.DB.Begin()
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

func (pdb *PostgresDb) BulkUpload(file [][]string) {
	var wg sync.WaitGroup
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
			pdb.SellOrCreate(&product)
			wg.Done()
			<-dbChan

		}(*product)
	}
	//wg.Wait()
}
