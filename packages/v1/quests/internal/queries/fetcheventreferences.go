package queries

import (
	"lib/utils"
	_types "quests/internal/types"
)

func FetchEventReferences(logger *utils.Logger) ([]string, error) {
	var response _types.PostHogQueryResponse
	var result []string

	err := PostHogQuery(`{"query":{"kind":"HogQLQuery","query":"select distinct events.event from events"}}`, &response)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	// extract the event names
	for _, value := range response.Results {
		result = append(result, value[0].(string))
	}

	return result, nil
}
