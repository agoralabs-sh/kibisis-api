package types

type ResponseBody struct {
	Account string       `json:"account,omitempty"`
	Error   interface{}  `json:"error,omitempty"`
	Quests  []DailyQuest `json:"quests,omitempty"`
}
