package config

import (
	"github.com/joho/godotenv"
	"log"
)

func NewConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
