package database

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
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
