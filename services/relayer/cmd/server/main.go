package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"lib/constants"
	"lib/utils"
	"os"
	"relayer/internal/database"
	_routes "relayer/internal/routes"
)

func main() {
	logger := utils.NewLogger()

	// setup database
	if err := database.InitDatabase(); err != nil {
		logger.Fatal(err)
	}

	// create server
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())

	// relay/pairing/create
	e.POST(fmt.Sprint(constants.RelayPath, constants.PairingPath, constants.CreatePath), _routes.NewPairingCreateRoute(logger))

	// start server
	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("RELAYER_PORT"))); err != nil {
		logger.Fatal(err)
	}
}
