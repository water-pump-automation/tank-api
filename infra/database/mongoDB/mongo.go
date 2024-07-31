package mongodb

import (
	"context"
	"tank-api/app/entity/tank"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitClient(ctx context.Context, connectionURI string) (client *mongo.Client, err error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
}

// Unused
func DoPing(ctx context.Context, client *mongo.Client) (err error) {
	return client.Ping(ctx, nil)
}

func NewCollection(ctx context.Context, client *mongo.Client, database, tankCollection, tankStateCollection string) tank.ITankDatabase {
	mongoDB := client.Database(database)

	tank := mongoDB.Collection(tankCollection)
	tankState := mongoDB.Collection(tankStateCollection)

	return NewTankMongoCollection(
		tank,
		NewTankStateMongoCollection(tankState),
	)
}
