package queries

import (
	"fmt"
	"lib/utils"
	_types "quests/internal/types"
)

func FetchDailyEvents(account string, logger *utils.Logger) ([]_types.DailyEvent, error) {
	var response _types.PostHogQueryResponse
	var result []_types.DailyEvent

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select events.event as \"event-name\", count() > 0 as \"completed\" from events where person.pdi.distinct_id = '%s' and events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo=' and events.timestamp >= today() group by \"event-name\" order by \"event-name\""}}`, account)
	err := PostHogQuery(query, &response)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	for _, value := range response.Results {
		result = append(result, _types.DailyEvent{
			Completed: int(value[1].(float64)) > 0,
			Name:      value[0].(string),
		})
	}

	return result, nil
}
