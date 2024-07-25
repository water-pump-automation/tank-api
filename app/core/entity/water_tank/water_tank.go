package water_tank

import (
	"time"
)

type Capacity float32

type State int

const (
	Empty State = iota
	Filling
	Full
)

type WaterTank struct {
	// IDs
	Name  string
	Group string

	// External updatable attributes
	CurrentWaterLevel Capacity

	// Internal updatable attributes
	TankState    State
	LastFullTime time.Time

	// Fixed attributes
	MaximumCapacity Capacity
}
