package models

import (
	"lib/models"
	"relayer/internal/types"
)

type Pairing struct {
	models.DefaultModel `bson:",inline"`
	Peers               []interface{} `json:"peers" bson:"peers"`
}

func NewPairing(client *types.ClientPeer, provider *types.ProviderPeer) *Pairing {
	peers := make([]interface{}, 0)

	if client != nil {
		peers = append(peers, client)
	}

	if provider != nil {
		peers = append(peers, provider)
	}

	return &Pairing{
		Peers: peers,
	}
}
