package types

// DailyQuest
// @Description The ID of the quest, whether the quest has been completed today and the total number of daily quests.
type DailyQuest struct {
	// The amount of times the quest has been completed
	Completed bool `json:"completed" example:"true"`
	// The ID of the quest
	Id string `json:"id" example:"send-native-currency-action"`
	// The total number of daily quests completed
	Total int `json:"total" example:"22"`
}
