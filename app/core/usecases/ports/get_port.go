package ports

import (
	"context"
	"water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
)

type IGetCapacity interface {
	GetMaximumCapacity(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (maximumCapacity water_tank.Capacity, err error_stack.Error)
}
