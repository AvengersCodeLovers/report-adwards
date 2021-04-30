package main

import (
	"log"
	"os"

	routes "github.com/AvengersCodeLovers/report-adwards/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := routes.SetupRouter()

	server.Run(":" + os.Getenv("APP_PORT"))
}
