package tank

import (
	"water-tank-api/core/entity/access"
	"water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
)

type Tank interface {
	GetData(tank string, group string) (MaximumCapacity water_tank.Capacity, accessToken access.AccessToken, err error_stack.ErrorStack)
}
