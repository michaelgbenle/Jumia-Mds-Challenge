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
	os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
}
