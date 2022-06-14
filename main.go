package main

import (
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
)

func main() {
	jumia := router.SetupRouter()
	err := jumia.Run(":2022")
	if err != nil {
		log.Fatal(err)
	}
}
