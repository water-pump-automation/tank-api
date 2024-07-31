package tank

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

type Tank struct {
	ID              string
	Name            string
	Group           string
	MaximumCapacity Capacity

	CurrentLevel Capacity
	LastFullTime *time.Time
}
