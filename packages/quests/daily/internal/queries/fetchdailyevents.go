package queries

import (
	"fmt"
	"lib/queries"
	"lib/types"
	"lib/utils"
	internaltypes "quests/internal/types"
)

func FetchDailyEvents(account string, logger *utils.Logger) ([]internaltypes.DailyEvent, error) {
	var response types.PostHogQueryResponse
	var result []internaltypes.DailyEvent

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select events.event as \"event-name\", count() as \"completed\" from events where person.pdi.distinct_id = '%s' and events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo=' and events.timestamp >= today() group by events.event, person.pdi.distinct_id order by count() desc"}}`, account)
	err := queries.PostHogQuery(query, &response)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	for _, value := range response.Results {
		result = append(result, internaltypes.DailyEvent{
			Amount: value[1].(int),
			Name:   value[0].(string),
		})
	}

	return result, nil
}
