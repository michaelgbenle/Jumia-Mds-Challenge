package main

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
)

func main() {

	fmt.Println("starting jumia app")
	jumia := router.SetupRouter()
	err := jumia.Run(":2022")
	if err != nil {
		log.Fatal(err)
	}

}
