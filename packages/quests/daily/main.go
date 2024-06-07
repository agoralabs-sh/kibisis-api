package main

import (
	"fmt"
	algosdktypes "github.com/algorand/go-algorand-sdk/v2/types"
	"golang.org/x/exp/slices"
	"lib/errors"
	"lib/queries"
	"lib/utils"
	"net/http"
	"os"
	_queries "quests/internal/queries"
	_types "quests/internal/types"
)

func Main(request _types.Request) *_types.Response {
	var dailyQuests []_types.DailyQuest

	logLevel := utils.LogLevelError

	switch os.Getenv("ENVIRONMENT") {
	case "development":
		logLevel = utils.LogLevelDebug
		break
	case "test":
		logLevel = utils.LogLevelSilent
		break
	default:
		logLevel = utils.LogLevelError
		break
	}

	logger := utils.NewLogger(logLevel)

	// only accept get requests
	if request.Http.Method != http.MethodGet {
		return &_types.Response{
			StatusCode: http.StatusMethodNotAllowed,
		}
	}

	_, err := algosdktypes.DecodeAddress(request.Account)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewInvalidAddressError(request.Account),
			},
			StatusCode: http.StatusBadRequest,
		}
	}

	logger.Debug("getting event references from posthog")

	eventReferences, err := queries.FetchEventReferences(logger)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch event references from posthog", err),
			},
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
		StatusCode: http.StatusOK,
	}
}
