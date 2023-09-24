package database_mock

import (
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankMockData struct {
}

func NewWaterTankMockData() *WaterTankMockData {
	return &WaterTankMockData{}
}

func (*WaterTankMockData) GetWaterTankState(names ...string) (state *data.WaterTankState, err stack.ErrorStack) {
	return
}

func (*WaterTankMockData) GetTankGroupState(groups ...string) (state []*data.WaterTankState, err stack.ErrorStack) {
	return
}

func (*WaterTankMockData) CreateWaterTank(name string, group string, capacity data.Capacity) (err stack.ErrorStack) {
	return
}

func (*WaterTankMockData) UpdateWaterTankState(name string, waterLevel data.Capacity) (state *data.WaterTankState, err stack.ErrorStack) {
	return
}
