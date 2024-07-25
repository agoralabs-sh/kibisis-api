package types

// LeaderboardResult
// @Description The account and the total daily quests the account has completed.
type LeaderboardResult struct {
	// The base 32 encoded address of the account
	Account string `json:"account" example:"TESTK4BURRDGVVHAX2FBY7CPRC2RTTVRRN4C2TVDCHRCXNTFGL3TVSDROE"`
	// The total number of daily quests completed
	Total int `json:"total" example:"22"`
}
