package queries

import (
	"fmt"
	"lib/types"
	"lib/utils"
	_types "quests/internal/types"
)

func FetchTotalDailyEvents(account string, logger *utils.Logger) ([]_types.TotalDailyEvent, error) {
	var response types.PostHogQueryResponse
	var result []_types.TotalDailyEvent

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select events.event as \"event-name\", count(distinct dateTrunc('day', events.timestamp)) AS \"total\" from events where person.pdi.distinct_id = '%s' and events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo=' group by \"event-name\" order by \"event-name\""}}`, account)
	err := utils.SendPostHogQuery(query, &response)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	for _, value := range response.Results {
		result = append(result, _types.TotalDailyEvent{
			Name:  value[0].(string),
			Total: int(value[1].(float64)),
		})
	}

	return result, nil
}
