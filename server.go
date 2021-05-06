package main

import (
	"fmt"
	"os"

	routes "github.com/AvengersCodeLovers/report-adwards/routes"
	"github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
)

func init() {
	f, err := os.OpenFile("storage/logs/logs.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	logrus.WithField("environment", os.Getenv("APP_ENV"))
	logrus.SetOutput(f)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	server := routes.SetupRouter()
	server.Run(":" + os.Getenv("APP_PORT"))
}
