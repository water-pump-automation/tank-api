package tank

import (
	"water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
)

type Tank interface {
	GetCapacity(tank string) (MaximumCapacity water_tank.Capacity, err error_stack.ErrorStack)
}
