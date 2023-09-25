package water_tank

import (
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water"
)

type State int

const (
	Empty State = iota
	Filling
	Full
)

type Capacity float32

type WaterTankState struct {
	Name              string
	Group             string
	MaximumCapacity   Capacity
	TankState         State
	CurrentWaterLevel Capacity
	Water             *water.Water
}

type WaterTankData interface {
	CreateWaterTank(name string, group string, capacity Capacity) (err stack.ErrorStack)
	UpdateWaterTankState(name string, waterLevel Capacity, levelState State) (state *WaterTankState, err stack.ErrorStack)
	GetWaterTankState(names ...string) (state *WaterTankState, err stack.ErrorStack)
	GetTankGroupState(groups ...string) (state []*WaterTankState, err stack.ErrorStack)
	GetAllTankGroupState() (state []*WaterTankState, err stack.ErrorStack)
}
