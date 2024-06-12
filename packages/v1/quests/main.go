package main

import (
	"fmt"
	algosdktypes "github.com/algorand/go-algorand-sdk/types"
	"golang.org/x/exp/slices"
	"lib/constants"
	"lib/errors"
	"lib/utils"
	"net/http"
	_queries "quests/internal/queries"
	_types "quests/internal/types"
)

// Main godoc
// @Summary Daily Quests
// @Description Gets the daily quests, as of 00:00:00 UTC, a given account.
// @Produce json
// @Param account query string true "account to get daily quests" Example(TESTK4BURRDGVVHAX2FBY7CPRC2RTTVRRN4C2TVDCHRCXNTFGL3TVSDROE)
// @Success 200 {object} _types.ResponseBody
// @Failure 400 "If the account param is not provided"
// @Failure 400 "If the account param is an invalid AVM address"
// @Failure 405 "If it is not a GET or OPTIONS request"
// @Failure 500
// @Header all {string} Cache-Control "public, max-age=3600"
// @Router /v1/quests [get]
func Main(request _types.Request) *_types.Response {
	var dailyQuests []_types.DailyQuest

	logger := utils.NewLogger()
	headers := _types.ResponseHeaders{
		CacheControl: fmt.Sprintf("public, max-age=%d", constants.HourInSeconds),
		ContentType:  "application/json",
	}

	// only accept get requests
	if request.Http.Method != http.MethodGet {
		return &_types.Response{
			Headers:    headers,
			StatusCode: http.StatusMethodNotAllowed,
		}
	}

	logger.Debug("validating params")

	if request.Account == "" {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewRequiredParamsError([]string{"account"}),
			},
			Headers:    headers,
			StatusCode: http.StatusBadRequest,
		}
	}

	logger.Debug(fmt.Sprintf("validating account \"%s\" param", request.Account))

	_, err := algosdktypes.DecodeAddress(request.Account)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewInvalidAddressError(request.Account),
			},
			Headers:    headers,
			StatusCode: http.StatusBadRequest,
		}
	}

	logger.Debug("getting event references from posthog")

	eventReferences, err := _queries.FetchEventReferences(logger)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch event references from posthog", err),
			},
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received the event references %s", eventReferences))

	logger.Debug(fmt.Sprintf("getting daily events from posthog for account \"%s\"", request.Account))

	dailyEvents, err := _queries.FetchDailyEvents(request.Account, logger)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch daily events from posthog", err),
			},
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received daily events from posthog for account \"%s\"", request.Account))

	// map the daily quests from teh events, defaulting to zero for quests that are not in the daily events from posthog
	for _, eventReference := range eventReferences {
		completed := 0
		index := slices.IndexFunc(dailyEvents, func(event _types.DailyEvent) bool {
			return event.Name == eventReference
		})

		// if the event reference is in the daily events, get the amount
		if index > -1 {
			completed = dailyEvents[index].Amount
		}

		dailyQuests = append(dailyQuests, _types.DailyQuest{
			Id:        eventReference,
			Completed: completed,
		})
	}

	return &_types.Response{
		Body: _types.ResponseBody{
			Account: request.Account,
			Quests:  dailyQuests,
		},
		Headers:    headers,
		StatusCode: http.StatusOK,
	}
}
