package queries

import (
	"fmt"
	"lib/types"
	"lib/utils"
	_constants "quests/internal/constants"
	_types "quests/internal/types"
)

// FetchLeaderboardResults Fetches a page of leaderboard results. The page is zero-indexed.
func FetchLeaderboardResults(page int, logger *utils.Logger) ([]_types.LeaderboardResult, error) {
	var response types.PostHogQueryResponse
	var results []_types.LeaderboardResult

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select person.pdi.distinct_id as \"account\", count(distinct dateTrunc('day', events.timestamp) || events.event) as \"total\" from events where events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo=' group by \"account\" order by \"total\" desc limit %d offset %d * %d"}}`, _constants.ResultLimit, page, _constants.ResultLimit)

	if err := utils.SendPostHogQuery(query, &response); err != nil {
		logger.Error(err)

		return nil, err
	}

	for _, value := range response.Results {
		results = append(results, _types.LeaderboardResult{
			Account: value[0].(string),
			Total:   int(value[1].(float64)),
		})
	}

	// if the results are nil, return an empty array
	if results == nil {
		return []_types.LeaderboardResult{}, nil
	}

	return results, nil
}
