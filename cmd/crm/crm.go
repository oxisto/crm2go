package main

import (
	"crm/db"
	"crm/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Starting...")

	db.InitPostgreSQL("localhost")

	router := routes.NewRouter()
	router.Run(":8100")
}
