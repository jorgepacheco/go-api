package main

import (
	"fmt"
	"github.com/go-first-steps/api"
	"github.com/go-first-steps/internal/logs"
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

	api.Start(applicationPort)

}
