package database

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"lib/utils"
	_enums "relayer/internal/enums"
	"relayer/internal/models"
)

func createPairingIndexes() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pairingTTLIndexName := "updated_at_ttl_index"

	pairingIndexes, err := mgm.Coll(&models.Pairing{}).Indexes().ListSpecifications(ctx)
	if err != nil {
		return err
	}

	// remove existing indexes
	for _, pairingIndex := range pairingIndexes {
		if pairingIndex.Name == pairingTTLIndexName {
			if _, err = mgm.Coll(&models.Pairing{}).Indexes().DropOne(ctx, pairingTTLIndexName); err != nil {
				return err
			}
		}
	}

	// create indexes
	if _, err = mgm.Coll(&models.Pairing{}).Indexes().CreateOne(ctx, utils.CreateTTLIndex(pairingTTLIndexName, "updated_at", int32(30), bson.D{
		{
			"$or",
			bson.A{
				bson.D{{
					"peers.status",
					bson.D{{"$eq", _enums.NotConnected}}, // remove if any of the peers are not connected
				}},
				bson.D{{
					"peer_count",
					bson.D{{"$lte", 1}}, // remove if there is only one peer
				}},
			},
		},
	})); err != nil {
		return err
	}

	return nil
}
