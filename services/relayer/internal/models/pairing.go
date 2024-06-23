package models

import (
	"github.com/fatih/structs"
	"lib/models"
	_types "relayer/internal/types"
)

type Pairing struct {
	models.DefaultModel `bson:",inline"`
	Peers               []map[string]interface{} `json:"peers" bson:"peers"`
	PeerCount           int                      `json:"-" bson:"peer_count"`
}

func NewPairing(client *_types.ClientPeer, provider *_types.ProviderPeer) *Pairing {
	peers := make([]map[string]interface{}, 0)

	if client != nil {
		peers = append(peers, structs.Map(client))
	}

	if provider != nil {
		peers = append(peers, structs.Map(provider))
	}

	return &Pairing{
		Peers:     peers,
		PeerCount: len(peers),
	}
}
