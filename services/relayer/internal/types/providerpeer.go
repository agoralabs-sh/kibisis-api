package types

type ProviderPeer struct {
	Info   ProviderInfo `json:"info" bson:"info"`
	Status string       `json:"status" bson:"status"`
	Type   string       `json:"type" bson:"type"`
}
