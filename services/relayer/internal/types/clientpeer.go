package types

type ClientPeer struct {
	Info   *ClientInfo `json:"info" bson:"info"`
	Status string      `json:"status" bson:"status"`
	Type   string      `json:"type" bson:"type"`
}
