package water_tank

import "water-tank-api/core/entity/water"

type State int

const (
	Empty State = iota
	Filling
	Full
)

type Capacity float32

type WaterTankState struct {
	Name              string
	MaximumCapacity   Capacity
	TankState         State
	CurrentWaterLevel Capacity
	Water             *water.Water
}

type WaterTankData interface {
	CreateWaterTank(name string, group string, capacity Capacity) (err error)
	UpdateWaterTankState(name string, waterLevel Capacity) (state *WaterTankState, err error)
	GetWaterTankState(names ...string) (state *WaterTankState, err error)
	GetTankGroupState(groups ...string) (state []*WaterTankState, err error)
}
