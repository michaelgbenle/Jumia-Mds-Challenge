package main

import (
	"github.com/michaelgbenle/jumiaMds/config"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
	"os"
)

func main() {
	config.NewConfig(".env")

	jumia := router.SetupRouter()
	port := os.Getenv("PORT")

	err := jumia.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
