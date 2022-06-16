package main

import (
	"github.com/michaelgbenle/jumiaMds/config"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
	"os"
)

func main() {
	config.NewConfig(".env")
	h := router.DataB()
	jumia := router.SetupRouter(h)
	port := os.Getenv("PORT")

	err := jumia.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
