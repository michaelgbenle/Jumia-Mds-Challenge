package router

import (
	"github.com/michaelgbenle/jumiaMds/database"
	"github.com/michaelgbenle/jumiaMds/handlers"
	"log"
	"os"
)

func DataB() handlers.Handler {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	Db := database.NewPostgresDb()
	err := Db.SetupDb(host, user, password, dbName, port)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.Handler{DB: Db}
	return handler

}
