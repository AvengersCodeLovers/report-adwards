package main

import (
	routes "github.com/hieudt-2054/report-adwards/routes"
)

func main() {
	server := routes.SetupRouter()

	server.Run(":8888")
}
