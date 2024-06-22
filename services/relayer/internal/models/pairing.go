package models

import (
	"lib/models"
	_types "relayer/internal/types"
)

type Pairing struct {
	models.DefaultModel `bson:",inline"`
	Peers               []interface{} `json:"peers" bson:"peers"`
	PeerCount           int           `json:"-" bson:"peer_count"`
}

func NewPairing(client *_types.ClientPeer, provider *_types.ProviderPeer) *Pairing {
	peers := make([]interface{}, 0)

	if client != nil {
		peers = append(peers, client)
	}

	if provider != nil {
		peers = append(peers, provider)
	}

	return &Pairing{
		Peers:     peers,
		PeerCount: len(peers),
	}
}
