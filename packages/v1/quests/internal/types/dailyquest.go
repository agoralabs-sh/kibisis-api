package types

// DailyQuest
// @Description The ID of the quest and the amount of times it has been completed.
type DailyQuest struct {
	// The amount of times the quest has been completed
	Completed int `json:"completed" example:"22"`
	// The ID of the quest
	Id string `json:"id" example:"send-native-currency-action"`
}
