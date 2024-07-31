package ports

import (
	"context"
	"tank-api/app/entity/tank"
)

const (
	INVALID_CAPACITY = -1
	EMPTY_CAPACITY   = 0
)

type IGetCapacity interface {
	GetMaximumCapacity(ctx context.Context, input *tank.GetTankStateInput) (maximumCapacity tank.Capacity, err error)
}

type ITankExists interface {
	Exists(ctx context.Context, input *tank.GetTankStateInput) (bool, error)
}
