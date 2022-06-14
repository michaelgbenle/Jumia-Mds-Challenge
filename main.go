package main

import (
	"fmt"
	"github.com/michaelgbenle/jumiaMds/router"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	fmt.Println("starting jumia app")
	jumia := router.SetupRouter(":2022", db)
}
