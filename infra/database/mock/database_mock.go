package database_mock

import (
	"context"
	"time"
	"water-tank-api/app/entity/water_tank"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPool struct{}

func (*MongoPool) Acquire() (water_tank.IConn, error) {
	return nil, nil
}
func (*MongoPool) AcquireTransaction() (water_tank.IConn, error) {
	return nil, nil
}

type MongoConn struct{}

func (*MongoConn) Release() error {
	return nil
}

func (*MongoConn) Query(ctx context.Context, callback water_tank.ConnCallback) {
	//
}

type WaterTankMockData struct {
	states map[string]map[string]*water_tank.WaterTank
}

var MockTimeNow = time.Now()

func NewWaterTankMockData(collection *mongo.Collection) *WaterTankMockData {
	return &WaterTankMockData{
		states: map[string]map[string]*water_tank.WaterTank{
			"GROUP_1": {
				"TANK_1": {
					Name:              "TANK_1",
					Group:             "GROUP_1",
					MaximumCapacity:   100,
					CurrentWaterLevel: 0,
					LastFullTime:      MockTimeNow,
				},
				"TANK_2": {
					Name:              "TANK_2",
					Group:             "GROUP_1",
					MaximumCapacity:   80,
					CurrentWaterLevel: 50,
					LastFullTime:      MockTimeNow,
				},
				"TANK_3": {
					Name:              "TANK_3",
					Group:             "GROUP_1",
					MaximumCapacity:   120,
					CurrentWaterLevel: 120,
					LastFullTime:      MockTimeNow,
				},
			},
			"GROUP_2": {
				"TANK_1": {
					Name:              "TANK_1",
					Group:             "GROUP_2",
					MaximumCapacity:   100,
					CurrentWaterLevel: 0,
					LastFullTime:      MockTimeNow,
				},
				"TANK_2": {
					Name:              "TANK_2",
					Group:             "GROUP_2",
					MaximumCapacity:   80,
					CurrentWaterLevel: 80,
					LastFullTime:      MockTimeNow,
				},
			},
			"GROUP_3": {
				"TANK_1": {
					Name:              "TANK_1",
					Group:             "GROUP_3",
					MaximumCapacity:   120,
					CurrentWaterLevel: 90,
					LastFullTime:      MockTimeNow,
				},
			},
		},
	}
}

func (tank *WaterTankMockData) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (state *water_tank.WaterTank, err error) {
	state = tank.states[input.Group][input.TankName]
	return
}

func (tank *WaterTankMockData) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanksInput) (state []*water_tank.WaterTank, err error) {
	if group, exists := tank.states[input.Group]; exists {
		for _, tank := range group {
			state = append(state, tank)
		}
	}
	return
}

func (tank *WaterTankMockData) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (state *water_tank.WaterTank, err error) {
	if _, exists := tank.states[input.Group]; !exists {
		tank.states[input.Group] = map[string]*water_tank.WaterTank{
			input.TankName: {
				Name:              input.TankName,
				Group:             input.Group,
				MaximumCapacity:   input.MaximumCapacity,
				CurrentWaterLevel: 0,
			},
		}
		return
	}

	tank.states[input.Group][input.TankName] = &water_tank.WaterTank{
		Name:              input.TankName,
		Group:             input.Group,
		MaximumCapacity:   input.MaximumCapacity,
		CurrentWaterLevel: 0,
	}

	return
}

func (tank *WaterTankMockData) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err error) {
	if group, exists := tank.states[input.TankName]; exists {
		group[input.TankName].CurrentWaterLevel = input.NewWaterLevel
	}
	return
}
