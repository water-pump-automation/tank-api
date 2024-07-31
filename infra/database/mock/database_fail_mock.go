package database_mock

import (
	"context"
	"errors"
	"tank-api/app/entity/tank"
)

type TankFailMockData struct{}

func NewTankFailMockData() *TankFailMockData {
	return &TankFailMockData{}
}

func (tank *TankFailMockData) GetTankState(ctx context.Context, input *tank.GetTankStateInput) (state *tank.Tank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *TankFailMockData) GetTankGroupState(ctx context.Context, input *tank.GetGroupTanksInput) (state []*tank.Tank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *TankFailMockData) CreateTank(ctx context.Context, input *tank.CreateInput) (tankState *tank.Tank, err error) {
	return nil, errors.New("ERROR")
}

func (tank *TankFailMockData) UpdateTankLevel(ctx context.Context, input *tank.UpdateLevelInput) (state *tank.Tank, err error) {
	return nil, errors.New("ERROR")
}
