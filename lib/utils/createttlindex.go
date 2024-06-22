package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTTLIndex Creates and index model that sets a field to expire after a lapsed time in seconds.
func CreateTTLIndex(indexName string, fieldName string, expireTimeInSeconds int32, filter interface{}) mongo.IndexModel {
	indexOptions := options.Index().SetName(indexName).SetExpireAfterSeconds(expireTimeInSeconds)

	if filter != nil {
		indexOptions.SetPartialFilterExpression(filter)
	}

	return mongo.IndexModel{
		Keys: bson.D{{
			fieldName,
			1,
		}},
		Options: indexOptions,
	}
}
