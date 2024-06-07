package queries

import (
	"fmt"
	"lib/queries"
	"lib/types"
	"lib/utils"
	_types "quests/internal/types"
)

func FetchDailyEvents(account string, logger *utils.Logger) ([]_types.DailyEvent, error) {
	var response types.PostHogQueryResponse
	var result []_types.DailyEvent

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select events.event as \"event-name\", count() as \"completed\" from events where person.pdi.distinct_id = '%s' and events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo=' and events.timestamp >= today() group by events.event, person.pdi.distinct_id order by count() desc"}}`, account)
	err := queries.PostHogQuery(query, &response)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	for _, value := range response.Results {
		result = append(result, _types.DailyEvent{
			Amount: int(value[1].(float64)),
			Name:   value[0].(string),
		})
	}

	return result, nil
}
