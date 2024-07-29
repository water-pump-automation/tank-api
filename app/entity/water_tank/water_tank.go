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
	Invalid
)

type WaterTank struct {
	Name            string
	Group           string
	MaximumCapacity Capacity

	CurrentWaterLevel Capacity
	LastFullTime      time.Time
}
