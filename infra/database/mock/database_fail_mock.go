package database_mock

import (
	"errors"
	"time"
	"water-tank-api/core/entity/access"
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankFailMockData struct{}

func NewWaterTankFailMockData() *WaterTankFailMockData {
	return &WaterTankFailMockData{}
}

func (tank *WaterTankFailMockData) GetWaterTankState(group string, names ...string) (state *data.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) GetTankGroupState(groups ...string) (state []*data.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) NotifyFullTank(name string, currentTime time.Time) (state *data.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity data.Capacity) (err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) UpdateWaterTankState(name string, waterLevel data.Capacity, levelState data.State) (state *data.WaterTank, err stack.ErrorStack) {
	err.AddEntityError(errors.New("ERROR"))
	return
}
