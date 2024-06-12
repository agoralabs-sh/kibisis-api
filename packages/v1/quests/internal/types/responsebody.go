package types

type ResponseBody struct {
	// The account address
	Account string      `json:"account,omitempty" example:"TESTK4BURRDGVVHAX2FBY7CPRC2RTTVRRN4C2TVDCHRCXNTFGL3TVSDROE"`
	Error   interface{} `json:"error,omitempty" swaggertype:"object"`
	// The completed quests
	Quests []DailyQuest `json:"quests,omitempty"`
}
