package database_mock

import (
	"context"
	"tank-api/app/entity/tank"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TankMockData struct {
	states map[string]map[string]*tank.Tank
}

func ptr(datetime time.Time) *time.Time {
	return &datetime
}

var MockTimeNow = ptr(time.Now())

func NewTankMockData(collection *mongo.Collection) *TankMockData {
	return &TankMockData{
		states: map[string]map[string]*tank.Tank{
			"GROUP_1": {
				"TANK_1": {
					Name:            "TANK_1",
					Group:           "GROUP_1",
					MaximumCapacity: 100,
					CurrentLevel:    0,
					LastFullTime:    MockTimeNow,
				},
				"TANK_2": {
					Name:            "TANK_2",
					Group:           "GROUP_1",
					MaximumCapacity: 80,
					CurrentLevel:    50,
					LastFullTime:    MockTimeNow,
				},
				"TANK_3": {
					Name:            "TANK_3",
					Group:           "GROUP_1",
					MaximumCapacity: 120,
					CurrentLevel:    120,
					LastFullTime:    MockTimeNow,
				},
			},
			"GROUP_2": {
				"TANK_1": {
					Name:            "TANK_1",
					Group:           "GROUP_2",
					MaximumCapacity: 100,
					CurrentLevel:    0,
					LastFullTime:    MockTimeNow,
				},
				"TANK_2": {
					Name:            "TANK_2",
					Group:           "GROUP_2",
					MaximumCapacity: 80,
					CurrentLevel:    80,
					LastFullTime:    MockTimeNow,
				},
			},
			"GROUP_3": {
				"TANK_1": {
					Name:            "TANK_1",
					Group:           "GROUP_3",
					MaximumCapacity: 120,
					CurrentLevel:    90,
					LastFullTime:    MockTimeNow,
				},
			},
		},
	}
}

func (data *TankMockData) GetTankState(ctx context.Context, input *tank.GetTankStateInput) (state *tank.Tank, err error) {
	state = data.states[input.Group][input.TankName]
	return
}

func (data *TankMockData) GetTankGroupState(ctx context.Context, input *tank.GetGroupTanksInput) (state []*tank.Tank, err error) {
	if group, exists := data.states[input.Group]; exists {
		for _, tank := range group {
			state = append(state, tank)
		}
	}
	return
}

func (data *TankMockData) CreateTank(ctx context.Context, input *tank.CreateInput) (state *tank.Tank, err error) {
	if _, exists := data.states[input.Group]; !exists {
		data.states[input.Group] = map[string]*tank.Tank{
			input.TankName: {
				Name:            input.TankName,
				Group:           input.Group,
				MaximumCapacity: input.MaximumCapacity,
				CurrentLevel:    0,
			},
		}
		return
	}

	data.states[input.Group][input.TankName] = &tank.Tank{
		Name:            input.TankName,
		Group:           input.Group,
		MaximumCapacity: input.MaximumCapacity,
		CurrentLevel:    0,
	}

	return
}

func (data *TankMockData) UpdateTankLevel(ctx context.Context, input *tank.UpdateLevelInput) (state *tank.Tank, err error) {
	if group, exists := data.states[input.TankName]; exists {
		group[input.TankName].CurrentLevel = input.NewLevel
	}
	return
}
