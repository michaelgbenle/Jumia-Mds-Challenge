package database

import (
	"github.com/michaelgbenle/jumiaMds/models"
	"log"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

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
			sellOrCreate(&product)
			wg.Done()
			<-dbChan

		}(*product)
	}
	//wg.Wait()
}
