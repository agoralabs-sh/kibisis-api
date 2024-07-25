package queries

import (
	"fmt"
	"lib/types"
	"lib/utils"
)

// FetchTotalAccountsCompleted Fetches the total number of accounts that have completed at least one event.
func FetchTotalAccountsCompleted(logger *utils.Logger) (int, error) {
	var response types.PostHogQueryResponse

	query := fmt.Sprintf(`{"query":{"kind":"HogQLQuery","query":"select count(distinct person.pdi.distinct_id) as \"total\" from events where events.properties.genesisHash = 'IXnoWtviVVJW5LGivNFc0Dq14V3kqaXuK2u5OQrdVZo='"}}`)
	if err := utils.SendPostHogQuery(query, &response); err != nil {
		logger.Error(err)

		return 0, err
	}

	return int(response.Results[0][0].(float64)), nil
}
