package database_mock

import data "water-tank-api/core/entity/water_tank"

type WaterTankMockData struct {
}

func NewWaterTankMockData() *WaterTankMockData {
	return &WaterTankMockData{}
}

func (*WaterTankMockData) GetWaterTankState(names ...string) (state *data.WaterTankState, err error) {
	return
}

func (*WaterTankMockData) GetTankGroupState(groups ...string) (state []*data.WaterTankState, err error) {
	return
}

func (*WaterTankMockData) CreateWaterTank(name string, group string, capacity data.Capacity) (err error) {
	return
}

func (*WaterTankMockData) UpdateWaterTankState(name string, waterLevel data.Capacity) (state *data.WaterTankState, err error) {
	return
}
