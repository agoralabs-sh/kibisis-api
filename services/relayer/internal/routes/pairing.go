package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"relayer/internal/models"
)

func NewPairingCreateRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		pairing := models.NewPairing(nil, nil)

		return c.JSON(http.StatusCreated, pairing)
	}
}
