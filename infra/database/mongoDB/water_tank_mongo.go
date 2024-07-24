package mongodb

import (
	"errors"
	"time"
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"

	"go.mongodb.org/mongo-driver/mongo"
)

type WaterTankMongoDB struct {
	collection *mongo.Collection
}

func (db *WaterTankMongoDB) CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity water_tank.Capacity) (err stack.ErrorStack) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return
}
func (db *WaterTankMongoDB) UpdateWaterTankState(name string, group string, waterLevel water_tank.Capacity, levelState water_tank.State) (state *water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return

}
func (db *WaterTankMongoDB) NotifyFullTank(name string, group string, currentTime time.Time) (state *water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return

}
func (db *WaterTankMongoDB) GetWaterTankState(group string, names ...string) (state *water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return

}
func (db *WaterTankMongoDB) GetTankGroupState(groups ...string) (state []*water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("NOT IMPLEMENTED"))
	return
}
