package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/jumiaMds/router"
	"log"
	"os"
)

func main() {
	fmt.Println("hello111")
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	jumia := router.SetupRouter()
	port := os.Getenv("PORT")
	fmt.Println("hello")
	err = jumia.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
