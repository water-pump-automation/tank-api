package database_mock

import (
	"errors"
	"time"
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankFailMockData struct{}

func NewWaterTankFailMockData() *WaterTankFailMockData {
	return &WaterTankFailMockData{}
}

func (tank *WaterTankFailMockData) GetWaterTankState(names ...string) (state *data.WaterTank, err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) GetTankGroupState(groups ...string) (state []*data.WaterTank, err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) NotifyFullTank(name string, currentTime time.Time) (state *data.WaterTank, err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) GetAllTankGroupState() (state []*data.WaterTank, err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) CreateWaterTank(name string, group string, capacity data.Capacity) (err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) UpdateWaterTankState(name string, waterLevel data.Capacity, levelState data.State) (state *data.WaterTank, err stack.ErrorStack) {
	err.AppendEntityError(errors.New("ERROR"))
	return
}
