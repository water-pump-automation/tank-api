package mongodb

import (
	"context"
	"tank-api/app/entity/tank"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type stateCollection struct {
	ID         string        `bson:"id"`
	TankID     string        `bson:"tank_id"`
	WaterLevel tank.Capacity `bson:"water_level"`
	Datetime   time.Time     `bson:"datetime"`
}

type TankStateMongoCollection struct {
	collection *mongo.Collection
}

func NewTankStateMongoCollection(collection *mongo.Collection) *TankStateMongoCollection {
	return &TankStateMongoCollection{
		collection: collection,
	}
}

func (db *TankStateMongoCollection) InsertState(ctx context.Context, tankID string, level tank.Capacity) (err error) {
	_, err = db.collection.InsertOne(
		ctx,
		bson.D{{Key: "tank_id", Value: tankID}, {Key: "level", Value: level}, {Key: "datetime", Value: time.Now()}},
	)

	return
}

func (db *TankStateMongoCollection) GetLastState(ctx context.Context, tankID string) (lastLevel tank.Capacity, err error) {
	var queryResult stateCollection

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "tank_id", Value: bson.D{{Key: "$eq", Value: tankID}}},
				},
			},
		},
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "datetime", Value: -1}})

	err = db.collection.FindOne(
		ctx,
		filter,
		opts,
	).Decode(&queryResult)

	if err != nil {
		return
	}

	return queryResult.WaterLevel, nil
}

func (db *TankStateMongoCollection) GetLastFullTime(ctx context.Context, tankID string, maximumCapacity tank.Capacity) (datetime *time.Time, err error) {
	queryResult := new(stateCollection)

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "tank_id", Value: bson.D{{Key: "$eq", Value: tankID}}},
				},
			},
		},
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "water_level", Value: bson.D{{Key: "$eq", Value: maximumCapacity}}},
				},
			},
		},
	}

	err = db.collection.FindOne(ctx, filter).Decode(&queryResult)

	if err != nil {
		return
	}

	return &queryResult.Datetime, nil
}
