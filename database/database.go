package database

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
)

//PostgresDb implements the DB interface
type PostgresDb struct {
	DB *gorm.DB
}

func NewPostgresDb() *PostgresDb {
	return &PostgresDb{}
}

func (pdb *PostgresDb) SetupDb(host, user, password, dbName, port string) error {

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

func (pdb *PostgresDb) GetProductSku(sku, country string) (*models.Product, error) {
	product := models.Product{}
	if err := pdb.DB.Where("sku= ? AND country=?", sku, country).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pdb *PostgresDb) SellStock(product *models.Product) (*models.Order, error) {
	purchaseStock := int(math.Abs(float64(product.Stock)))

	err := pdb.DB.First(product, "sku=? AND country = ? AND stock >= ?", product.Sku, product.Country, purchaseStock).Error
	if product.ID <= 0 {

		return nil, err
	}
	order := models.Order{
		ProductId: product.ID,
		Quantity:  uint(purchaseStock),
	}

	//create the order
	if err = pdb.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	//Update product amount to reflect change
	if err = pdb.DB.Model(models.Product{}).Where("id = ?", product.ID).
		Updates(models.Product{Stock: product.Stock - purchaseStock}).Error; err != nil {
		return nil, err
	}

	return &order, nil
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
			log.Println("incomplete data", i)
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
func (pdb *PostgresDb) SellOrCreate(product *models.Product) {
	if int(product.Stock) < 0 {
		_, err := pdb.SellStock(product)
		if err != nil {
			return
		}
	} else {
		pdb.ProductCreate(product)
	}
}
func (pdb *PostgresDb) ProductCreate(product *models.Product) *models.Product {
	changeInStock := product.Stock
	trans := pdb.DB.Begin()
	var dbProduct models.Product
	trans.Where("sku =? AND country = ?", product.Sku, product.Country).First(dbProduct) // needs review

	if dbProduct.ID == 0 {
		if err := trans.Create(dbProduct).Error; err != nil {
			trans.Rollback()
			return &dbProduct
		}
		trans.Commit()
		return product
	}
	//update product to reflect change
	if err := trans.Model(models.Product{}).
		Where("id = ?", dbProduct.ID).Update("stock", changeInStock+dbProduct.Stock).Error; err != nil {
		trans.Rollback()
		return &models.Product{}
	}
	//product.Stock += changeInStock
	trans.Commit()
	return &dbProduct
}
