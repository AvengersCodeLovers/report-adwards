package main

import (
	routes "github.com/AvengersCodeLovers/report-adwards/routes"
)

func main() {
	server := routes.SetupRouter()

	server.Run(":8888")
}
