package database_mock

import (
	"context"
	"errors"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
)

type WaterTankFailMockData struct{}

func NewWaterTankFailMockData() *WaterTankFailMockData {
	return &WaterTankFailMockData{}
}

func (tank *WaterTankFailMockData) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (state *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (state []*water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (tankState *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("ERROR"))
	return
}

func (tank *WaterTankFailMockData) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err stack.Error) {
	err.AddEntityError(errors.New("ERROR"))
	return
}
