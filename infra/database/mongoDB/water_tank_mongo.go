package mongodb

import (
	"context"
	"errors"
	"water-tank-api/app/entity/water_tank"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPool struct{}

func (*MongoPool) Acquire() (water_tank.IConn, error) {
	return nil, nil
}
func (*MongoPool) AcquireTransaction() (water_tank.IConn, error) {
	return nil, nil
}

type MongoConn struct{}

func (*MongoConn) Release() error {
	return nil
}

func (*MongoConn) Query(ctx context.Context, callback water_tank.ConnCallback) {
	//
}

type WaterTankMongoDB struct {
	collection *mongo.Collection
}

func NewWaterTankMongoDB(collection *mongo.Collection) *WaterTankMongoDB {
	return &WaterTankMongoDB{
		collection: collection,
	}
}

func (db *WaterTankMongoDB) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (state *water_tank.WaterTank, err error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

func (db *WaterTankMongoDB) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err error) {
	return nil, errors.New("NOT IMPLEMENTED")

}

func (db *WaterTankMongoDB) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (state *water_tank.WaterTank, err error) {
	return nil, errors.New("NOT IMPLEMENTED")

}

func (db *WaterTankMongoDB) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanksInput) (state []*water_tank.WaterTank, err error) {
	return nil, errors.New("NOT IMPLEMENTED")
}
