package main

import (
	"fmt"
	"github.com/go-first-steps/api"
	"github.com/go-first-steps/internal/data"
	"github.com/go-first-steps/internal/logs"
	"log"
	"os"
)

const defaultPort = "8080"

func main() {

	fmt.Println("::: Initial go-first-steps")

	applicationPort := os.Getenv("PORT")

	logs.InitDefault("dev") //hardcoded at dev environment for the PoC API

	if applicationPort == "" {
		applicationPort = defaultPort
	}
	logs.Log().Info(applicationPort)

	fmt.Println("::: Connect to DataBase")

	initDb()

	api.Start(applicationPort)

}

func initDb() {

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

}
