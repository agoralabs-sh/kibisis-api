package routes

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"lib/enums"
	"lib/errors"
	"lib/utils"
	"net/http"
	_enums "relayer/internal/enums"
	_models "relayer/internal/models"
	_requests "relayer/internal/requests"
	_types "relayer/internal/types"
)

func NewPairingCreateRoute(logger *utils.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody _requests.PairingCreateRequestBody
		var validationError *errors.ValidationError
		var unknownError *errors.UnknownError

		// parse the request body
		if err := c.Bind(&requestBody); err != nil {
			unknownError = errors.NewUnknownError("failed to parse the request body", err)

			logger.Error(unknownError)

			return c.JSON(http.StatusInternalServerError, unknownError)
		}
		// validate the request body
		if validationErrors := requestBody.Validate(); validationErrors != nil {
			validationError = errors.NewValidationError(enums.BodyFieldType, validationErrors)

			logger.Error(validationError)

			return c.JSON(http.StatusBadRequest, validationErrors)
		}

		pairing := _models.NewPairing(&_types.ClientPeer{
			Info: &_types.ClientInfo{
				Description: requestBody.Description,
				ID:          uuid.New().String(),
				Host:        requestBody.Host,
				Icon:        requestBody.Icon,
				Name:        requestBody.Name,
			},
			Status: _enums.NotConnected,
			Type:   _enums.ClientType,
		}, nil)

		// save to the database
		if err := mgm.Coll(pairing).Create(pairing); err != nil {
			unknownError = errors.NewUnknownError("failed to save pairing", err)

			logger.Error(unknownError)

			return c.JSON(http.StatusInternalServerError, unknownError)
		}

		return c.JSON(http.StatusCreated, pairing)
	}
}
