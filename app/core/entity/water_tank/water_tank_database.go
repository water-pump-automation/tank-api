package water_tank

import (
	"time"
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
)

type State int

const (
	Empty State = iota
	Filling
	Full
)

type Capacity float32

type WaterTank struct {
	// IDs
	Name   string
	Group  string
	Access access.AccessToken

	// External updatable attributes
	CurrentWaterLevel Capacity

	// Internal updatable attributes
	TankState    State
	LastFullTime time.Time

	// Fixed attributes
	MaximumCapacity Capacity
}

type WaterTankData interface {
	CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity Capacity) (err stack.ErrorStack)
	UpdateWaterTankState(name string, group string, waterLevel Capacity, levelState State) (state *WaterTank, err stack.ErrorStack)
	NotifyFullTank(name string, group string, currentTime time.Time) (state *WaterTank, err stack.ErrorStack)
	GetWaterTankState(group string, names ...string) (state *WaterTank, err stack.ErrorStack)
	GetTankGroupState(groups ...string) (state []*WaterTank, err stack.ErrorStack)
}
