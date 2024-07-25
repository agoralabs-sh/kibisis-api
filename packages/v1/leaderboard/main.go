package main

import (
	"fmt"
	"lib/constants"
	"lib/errors"
	"lib/utils"
	"math"
	"net/http"
	_constants "quests/internal/constants"
	_queries "quests/internal/queries"
	_types "quests/internal/types"
	"strconv"
)

// Main godoc
// @Summary Quests Leaderboard
// @Description Gets the quests' leaderboard for all accounts.
// @Produce json
// @Param page query string false "the page of results, the first page starts at 0" Example(0)
// @Success 200 {object} _types.ResponseBody
// @Failure 405 "If it is not a GET request"
// @Failure 500
// @Header all {string} Cache-Control "public, max-age=3600"
// @Router /v1/leaderboard [get]
func Main(request _types.Request) *_types.Response {
	var err error
	var page = 0

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

	if request.Page != "" {
		if page, err = strconv.Atoi(request.Page); err != nil {
			logger.Debug(fmt.Sprintf("failed to parse page \"%s\" to integer", request.Page))
		}
	}

	logger.Debug("fetching total number of accounts that have completed at least one event")

	total, err := _queries.FetchTotalAccountsCompleted(logger)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewPostHogError("failed to fetch the total amount of accounts that have completed an event from posthog", err),
			},
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received the total amount of accounts that have completed an event: %d", total))

	pageCount := int(math.Ceil(float64(total / _constants.ResultLimit))) // pages are zero-based

	logger.Debug(fmt.Sprintf("fetching total daily event for each account from page %d of %d from posthog", page, pageCount))

	results, err := _queries.FetchLeaderboardResults(page, logger)
	if err != nil {
		return &_types.Response{
			Body: _types.ResponseBody{
				Error: errors.NewPostHogError(fmt.Sprintf("failed to fetch the total number of completed events for page %d of %d from posthog", page, pageCount), err),
			},
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}
	}

	logger.Debug(fmt.Sprintf("received total number of completed events for page %d of %d from posthog", page, pageCount))

	return &_types.Response{
		Body: _types.ResponseBody{
			Metadata: _types.PaginationMetadata{
				Page:      page,
				PageCount: pageCount,
				PageSize:  _constants.ResultLimit,
				Total:     total,
			},
			Results: results,
		},
		Headers:    headers,
		StatusCode: http.StatusOK,
	}
}
