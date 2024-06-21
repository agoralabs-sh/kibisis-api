package main

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib/constants"
	"lib/utils"
	"os"
	"relayer/internal/routes"
)

func main() {
	logger := utils.NewLogger()

	// connection
	connectionOptions := options.Client().ApplyURI(fmt.Sprintf("%s%s:%s@%s:%s", os.Getenv("MONGODB_PROTOCOL"), os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT")))
	if err := mgm.SetDefaultConfig(nil, os.Getenv("MONGODB_NAME"), connectionOptions); err != nil {
		logger.Fatal(err)
	}

	// create server
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())

	// /pairing/create
	e.POST(fmt.Sprint(constants.RelayPath, constants.PairingPath, constants.CreatePath), routes.NewPairingCreateRoute())

	// start server
	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("RELAYER_PORT"))); err != nil {
		logger.Fatal(err)
	}
}
