package main

import (
	"os"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/data"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server"
	"github.com/diegoclair/go_utils-lib/logger"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/service"
)

func main() {
	logger.Info("Reading the initial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}

	svc := service.New(db)
	server := server.InitServer(svc)
	logger.Info("About to start the application...")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	if err := server.Start(":" + port); err != nil {
		panic(err)
	}
}
