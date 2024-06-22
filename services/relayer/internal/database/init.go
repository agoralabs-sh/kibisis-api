package database

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func InitDatabase() error {
	connectionOptions := options.Client().ApplyURI(fmt.Sprintf("%s%s:%s@%s:%s", os.Getenv("MONGODB_PROTOCOL"), os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT")))

	if err := mgm.SetDefaultConfig(nil, os.Getenv("MONGODB_NAME"), connectionOptions); err != nil {
		return err
	}

	// create indexes
	if err := createPairingIndexes(); err != nil {
		return err
	}

	return nil
}
