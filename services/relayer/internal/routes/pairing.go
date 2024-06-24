package routes

import (
	"fmt"
	"github.com/fatih/structs"
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

func NewPairingAddRoute(logger *utils.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody _requests.PairingAddRequestBody
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

			logger.Debug(validationError)

			return c.JSON(http.StatusBadRequest, validationErrors)
		}

		pairingId := c.Param("id")
		pairing := &_models.Pairing{}

		// get the pairing
		if err := mgm.Coll(pairing).FindByID(pairingId, pairing); err != nil {
			logger.Error(err)

			return c.JSON(http.StatusNotFound, errors.NewPairingNotFoundError(pairingId))
		}

		if providerPeer := pairing.GetProviderPeer(); providerPeer != nil {
			logger.Debug(fmt.Sprintf("peer \"%s\" already added to pairing \"%s\"", providerPeer.Info.Name, pairingId))

			return c.JSON(http.StatusForbidden, fmt.Sprintf("provider peer already added to pairing \"%s\"", pairingId))
		}

		// add the provider
		pairing.Peers = append(pairing.Peers, structs.Map(&_types.ProviderPeer{
			Info: &_types.ProviderInfo{
				ID:   requestBody.ID,
				Name: requestBody.Name,
			},
			Status: _enums.NotConnected,
			Type:   _enums.ProviderType,
		}))

		// save to the database
		if err := mgm.Coll(pairing).Update(pairing); err != nil {
			unknownError = errors.NewUnknownError(fmt.Sprintf("failed to save to pairing \"%s\"", pairingId), err)

			logger.Error(unknownError)

			return c.JSON(http.StatusInternalServerError, unknownError)
		}

		return c.JSON(http.StatusOK, pairing)
	}
}

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
