package mongodb

import (
	"context"
	"errors"
	"tank-api/app/entity/tank"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var empty = tank.Capacity(0.0)

type tankCollection struct {
	ID              string        `bson:"id"`
	Name            string        `bson:"name"`
	Group           string        `bson:"group"`
	MaximumCapacity tank.Capacity `bson:"maximum_capacity"`
}

type TankMongoCollection struct {
	collection      *mongo.Collection
	stateCollection *TankStateMongoCollection
}

func NewTankMongoCollection(collection *mongo.Collection, stateCollection *TankStateMongoCollection) *TankMongoCollection {
	return &TankMongoCollection{
		collection:      collection,
		stateCollection: stateCollection,
	}
}

func (db *TankMongoCollection) CreateTank(ctx context.Context, input *tank.CreateInput) (state *tank.Tank, err error) {
	state = new(tank.Tank)

	tankID, err := db.collection.InsertOne(
		ctx,
		bson.D{
			{Key: "name", Value: input.TankName},
			{Key: "group", Value: input.Group},
			{Key: "maximum_capacity", Value: input.MaximumCapacity},
		},
	)

	if err != nil {
		return nil, err
	}

	err = db.stateCollection.InsertState(ctx, tankID.InsertedID.(string), empty)

	if err != nil {
		return nil, err
	}

	state.Name = input.TankName
	state.Group = input.Group
	state.MaximumCapacity = input.MaximumCapacity

	state.CurrentLevel = empty

	return
}

func (db *TankMongoCollection) UpdateTankLevel(ctx context.Context, input *tank.UpdateLevelInput) (state *tank.Tank, err error) {
	state, err = db.GetTankState(ctx, &tank.GetTankStateInput{
		TankName: input.TankName,
		Group:    input.Group,
	})

	if err != nil || state == nil {
		return
	}

	err = db.stateCollection.InsertState(ctx, state.ID, input.NewLevel)

	if err != nil {
		return nil, err
	}

	return
}

func (db *TankMongoCollection) GetTankState(ctx context.Context, input *tank.GetTankStateInput) (state *tank.Tank, err error) {
	queryResult := new(tankCollection)
	state = new(tank.Tank)

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "name", Value: bson.D{{Key: "$eq", Value: input.TankName}}},
				},
			},
		},
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "group", Value: bson.D{{Key: "$eq", Value: input.Group}}},
				},
			},
		},
	}

	err = db.collection.FindOne(ctx, filter).Decode(queryResult)
	if err != nil {
		return
	}

	state.Name = queryResult.Name
	state.Group = queryResult.Group
	state.MaximumCapacity = queryResult.MaximumCapacity

	state.CurrentLevel, err = db.stateCollection.GetLastState(ctx, queryResult.ID)

	if err != nil {
		return nil, err
	}

	state.LastFullTime, err = db.stateCollection.GetLastFullTime(ctx, queryResult.ID, queryResult.MaximumCapacity)

	if err != nil {
		return nil, err
	}

	return
}

func (db *TankMongoCollection) GetTankGroupState(ctx context.Context, input *tank.GetGroupTanksInput) (state []*tank.Tank, err error) {
	var results []bson.M

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "group", Value: bson.D{{Key: "$eq", Value: input.Group}}},
				},
			},
		},
	}

	cursor, err := db.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error retrieving tank state")
		}
	}()

	for _, result := range results {
		tankState, err := db.GetTankState(ctx, &tank.GetTankStateInput{
			TankName: result["name"].(string),
			Group:    result["group"].(string),
		})

		if err != nil {
			return nil, err
		}
		state = append(state, tankState)
	}

	return
}
