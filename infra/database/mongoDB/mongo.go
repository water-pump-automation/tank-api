package mongodb

import (
	"context"
	"water-tank-api/app/entity/water_tank"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitClient(ctx context.Context, connectionURI string) (client *mongo.Client, err error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
}

func DoPing(ctx context.Context, client *mongo.Client) (err error) {
	return client.Ping(ctx, nil)
}

func NewCollection(ctx context.Context, client *mongo.Client, database, collection string) water_tank.IWaterTankDatabase {
	mongoCollection := client.Database(database).Collection(collection)

	return NewWaterTankMongoDB(mongoCollection)
}
