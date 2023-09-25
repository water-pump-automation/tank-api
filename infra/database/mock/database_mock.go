package database_mock

import (
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankMockData struct {
	states map[string]*data.WaterTankState
}

func NewWaterTankMockData() *WaterTankMockData {
	return &WaterTankMockData{
		states: map[string]*data.WaterTankState{
			"TANK_1": {
				Name:              "TANK_1",
				Group:             "GROUP_1",
				MaximumCapacity:   100,
				TankState:         data.Empty,
				CurrentWaterLevel: 0,
				Water:             nil,
			},
			"TANK_2": {
				Name:              "TANK_2",
				Group:             "GROUP_1",
				MaximumCapacity:   80,
				TankState:         data.Filling,
				CurrentWaterLevel: 50,
				Water:             nil,
			},
			"TANK_3": {
				Name:              "TANK_3",
				Group:             "GROUP_1",
				MaximumCapacity:   120,
				TankState:         data.Full,
				CurrentWaterLevel: 120,
				Water:             nil,
			},
			"TANK_4": {
				Name:              "TANK_4",
				Group:             "GROUP_2",
				MaximumCapacity:   100,
				TankState:         data.Empty,
				CurrentWaterLevel: 0,
				Water:             nil,
			},
			"TANK_5": {
				Name:              "TANK_5",
				Group:             "GROUP_2",
				MaximumCapacity:   80,
				TankState:         data.Full,
				CurrentWaterLevel: 80,
				Water:             nil,
			},
			"TANK_6": {
				Name:              "TANK_6",
				Group:             "GROUP_3",
				MaximumCapacity:   120,
				TankState:         data.Filling,
				CurrentWaterLevel: 90,
				Water:             nil,
			},
		},
	}
}

func (tank *WaterTankMockData) GetWaterTankState(names ...string) (state *data.WaterTankState, err stack.ErrorStack) {
	state = tank.states[names[0]]
	return
}

func (tank *WaterTankMockData) GetTankGroupState(groups ...string) (state []*data.WaterTankState, err stack.ErrorStack) {
	for _, tank := range tank.states {
		if tank.Group == groups[0] {
			state = append(state, tank)
		}
	}
	return
}

func (tank *WaterTankMockData) GetAllTankGroupState() (state []*data.WaterTankState, err stack.ErrorStack) {
	for _, tank := range tank.states {
		state = append(state, tank)
	}

	return
}

func (tank *WaterTankMockData) CreateWaterTank(name string, group string, capacity data.Capacity) (err stack.ErrorStack) {
	tank.states[name] = &data.WaterTankState{
		Name:              name,
		Group:             group,
		MaximumCapacity:   capacity,
		TankState:         data.Empty,
		CurrentWaterLevel: 0,
		Water:             nil,
	}
	return
}

func (tank *WaterTankMockData) UpdateWaterTankState(name string, waterLevel data.Capacity, levelState data.State) (state *data.WaterTankState, err stack.ErrorStack) {
	if tank, exists := tank.states[name]; exists {
		tank.CurrentWaterLevel = waterLevel
		tank.TankState = levelState
	}
	return
}
