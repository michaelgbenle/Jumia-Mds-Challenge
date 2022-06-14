package main

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
	"os"
)

func main() {
	jumia := router.SetupRouter()
	port := os.Getenv("PORT")
	err := jumia.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
}
