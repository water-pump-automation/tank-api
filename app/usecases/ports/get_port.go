package ports

import (
	"context"
	"water-tank-api/app/entity/water_tank"
)

const (
	INVALID_CAPACITY = -1
	EMPTY_CAPACITY   = 0
)

type IGetCapacity interface {
	GetMaximumCapacity(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (maximumCapacity water_tank.Capacity, err error)
}

type ITankExists interface {
	Exists(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (bool, error)
}
