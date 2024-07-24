package mongodb

import (
	"context"
	"water-tank-api/app/core/entity/water_tank"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitClient(ctx context.Context, connectionURI string) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	return
}

func NewCollection(ctx context.Context, client *mongo.Client, database, collection string) water_tank.WaterTankData {
	mongoCollection := client.Database(database).Collection(collection)

	return &WaterTankMongoDB{
		collection: mongoCollection,
	}
}
