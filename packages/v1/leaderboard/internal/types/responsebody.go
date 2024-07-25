package types

type ResponseBody struct {
	Error interface{} `json:"error,omitempty" swaggertype:"object"`
	// The pagination metadata
	Metadata PaginationMetadata `json:"metadata"`
	// The leaderboard results; the completed quests for each account
	Results []LeaderboardResult `json:"results"`
}
