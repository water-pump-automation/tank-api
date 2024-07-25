package mongodb

import (
	"context"
	"errors"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"

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

func (db *WaterTankMongoDB) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (state *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return
}

func (db *WaterTankMongoDB) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return

}

func (db *WaterTankMongoDB) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (state *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return

}

func (db *WaterTankMongoDB) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (state []*water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return
}
