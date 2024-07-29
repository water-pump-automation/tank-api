package database_mock

import (
	"context"
	"errors"
	"water-tank-api/app/entity/water_tank"
)

type WaterTankFailMockData struct{}

func NewWaterTankFailMockData() *WaterTankFailMockData {
	return &WaterTankFailMockData{}
}

func (tank *WaterTankFailMockData) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (state *water_tank.WaterTank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *WaterTankFailMockData) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanksInput) (state []*water_tank.WaterTank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *WaterTankFailMockData) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (tankState *water_tank.WaterTank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *WaterTankFailMockData) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err error) {
	return nil, errors.New("ERROR")
}
