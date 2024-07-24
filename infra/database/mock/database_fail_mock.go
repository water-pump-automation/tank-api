package database_mock

import (
	"errors"
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
)

type WaterTankFailMockData struct{}

func NewWaterTankFailMockData() *WaterTankFailMockData {
	return &WaterTankFailMockData{}
}

func (tank *WaterTankFailMockData) GetWaterTankState(group string, names ...string) (state *water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) GetTankGroupState(groups ...string) (state []*water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity water_tank.Capacity) (err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) UpdateWaterTankState(name string, group string, waterLevel water_tank.Capacity, levelState water_tank.State) (state *water_tank.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}
