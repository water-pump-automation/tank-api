package water_tank

import (
	"time"
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

type WaterTank struct {
	Name              string
	Group             string
	MaximumCapacity   Capacity
	TankState         State
	CurrentWaterLevel Capacity
	LastFullTime      time.Time
	Water             *water.Water
}

type WaterTankData interface {
	CreateWaterTank(name string, group string, capacity Capacity) (err stack.ErrorStack)
	UpdateWaterTankState(name string, waterLevel Capacity, levelState State) (state *WaterTank, err stack.ErrorStack)
	NotifyFullTank(name string, currentTime time.Time) (state *WaterTank, err stack.ErrorStack)
	GetWaterTankState(names ...string) (state *WaterTank, err stack.ErrorStack)
	GetTankGroupState(groups ...string) (state []*WaterTank, err stack.ErrorStack)
	GetAllTankGroupState() (state []*WaterTank, err stack.ErrorStack)
}
