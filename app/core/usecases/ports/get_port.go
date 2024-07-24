package ports

import (
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
)

type GetUsecase interface {
	GetData(tank string, group string) (MaximumCapacity water_tank.Capacity, accessToken access.AccessToken, err error_stack.ErrorStack)
}
