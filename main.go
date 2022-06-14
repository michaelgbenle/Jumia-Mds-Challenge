package main

import (
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	jumia := router.SetupRouter()
	port := os.Getenv("PORT")

	err = jumia.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
