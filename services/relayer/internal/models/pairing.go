package models

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"lib/models"
	_enums "relayer/internal/enums"
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

func (p Pairing) GetClientPeer() *_types.ClientPeer {
	var result *_types.ClientPeer

	for _, peer := range p.Peers {
		if peer["Type "] == _enums.ClientType {
			if err := mapstructure.Decode(peer, &result); err != nil {
				return nil
			}

			return result
		}
	}

	return nil
}

func (p Pairing) GetProviderPeer() *_types.ProviderPeer {
	var result *_types.ProviderPeer

	for _, peer := range p.Peers {
		if peer["Type "] == _enums.ProviderType {
			if err := mapstructure.Decode(peer, &result); err != nil {
				return nil
			}

			return result
		}
	}

	return nil
}
