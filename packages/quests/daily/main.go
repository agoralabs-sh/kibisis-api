package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"lib/errors"
	"lib/queries"
	"lib/utils"
	"net/http"
	"os"
	internalqueries "quests/internal/queries"
	internaltypes "quests/internal/types"
)

func Main(request internaltypes.Request) *internaltypes.Response {
	var dailyQuests []internaltypes.DailyQuest

	logLevel := utils.LogLevelError

	if os.Getenv("ENVIRONMENT") == "development" {
		logLevel = utils.LogLevelDebug
	}

	logger := utils.NewLogger(logLevel)

	// only accept get requests
	//if request.Http.Method != http.MethodGet {
	//	return &internaltypes.Response{
	//		StatusCode: http.StatusMethodNotAllowed,
	//	}
	//}

	logger.Debug("getting event references from posthog")

	eventReferences, err := queries.FetchEventReferences(logger)
	if err != nil {
		return &internaltypes.Response{
			Body: internaltypes.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch event references from posthog", err),
			},
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received the event references %s", eventReferences))

	logger.Debug(fmt.Sprintf("getting daily events from posthog for account \"%s\"", request.Account))

	dailyEvents, err := internalqueries.FetchDailyEvents(request.Account, logger)
	if err != nil {
		return &internaltypes.Response{
			Body: internaltypes.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch daily events from posthog", err),
			},
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received daily events from posthog for account \"%s\"", request.Account))

	// map the daily quests from teh events, defaulting to zero for quests that are not in the daily events from posthog
	for _, eventReference := range eventReferences {
		completed := 0
		index := slices.IndexFunc(dailyEvents, func(event internaltypes.DailyEvent) bool {
			return event.Name == eventReference
		})

		// if the event reference is in the daily events, get the amount
		if index > -1 {
			completed = dailyEvents[index].Amount
		}

		dailyQuests = append(dailyQuests, internaltypes.DailyQuest{
			Id:        eventReference,
			Completed: completed,
		})
	}

	return &internaltypes.Response{
		Body: internaltypes.ResponseBody{
			Account: request.Account,
			Quests:  dailyQuests,
		},
		StatusCode: http.StatusOK,
	}
}
