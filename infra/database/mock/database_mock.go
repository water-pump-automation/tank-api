package database_mock

import "water-tank-api/core/entity/data"

type WaterTankMockData struct {
}

func NewWaterTankMockData() *WaterTankMockData {
	return &WaterTankMockData{}
}

func (*WaterTankMockData) GetDataByName(names ...string) (state *data.WaterTankState, err error) {
	return
}

func (*WaterTankMockData) GetDataByGroup(groups ...string) (state []*data.WaterTankState, err error) {
	return
}

func (*WaterTankMockData) CreateWaterTank(name string, group string) (err error) {
	return
}

func (*WaterTankMockData) UpdateWaterTankState(name string, waterLevel *data.Capacity) (state *data.WaterTankState, err error) {
	return
}
