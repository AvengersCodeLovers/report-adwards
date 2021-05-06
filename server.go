package main

import (
	"github.com/AvengersCodeLovers/report-adwards/util"

	routes "github.com/AvengersCodeLovers/report-adwards/routes"
)

func main() {
	util.LoadEnvVars()
	util.UseJSONLogFormat()

	server := routes.SetupRouter()

	server.Run(":" + util.GetEnv("APP_PORT", "8888"))
}
