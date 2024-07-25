package update_tank_state

import (
	"context"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type UpdateWaterTank struct {
	tank       water_tank.IWaterTankDatabase
	getUsecase ports.IGetCapacity
}

func NewWaterTankUpdate(tank water_tank.IWaterTankDatabase, getUsecase ports.IGetCapacity) *UpdateWaterTank {
	return &UpdateWaterTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *UpdateWaterTank) Update(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (err stack.Error) {
	var maximumCapacity water_tank.Capacity

	maximumCapacity, err = conn.getUsecase.GetMaximumCapacity(ctx, connection, &water_tank.GetWaterTankState{
		TankName: input.TankName,
		Group:    input.Group,
	})

	if err.HasError() {
		if entity := err.EntityError(); entity != nil {
			err.AppendUsecaseError(ErrWaterTankErrorServerError(entity.Error()))
			return
		}

		err.AppendUsecaseError(ErrWaterTankErrorNotFound(input.TankName))
		return
	}

	if input.NewWaterLevel < 0 {
		err.AppendUsecaseError(ErrWaterTankCurrentWaterLevelSmallerThanZero)
		return
	}

	if input.NewWaterLevel > maximumCapacity {
		err.AppendUsecaseError(ErrWaterTankCurrentWaterLevelBiggerThanMax)
		return
	} else if input.NewWaterLevel == maximumCapacity {
		input.State = water_tank.Full
	} else if input.NewWaterLevel == 0 {
		input.State = water_tank.Empty
	} else if input.NewWaterLevel < maximumCapacity {
		input.State = water_tank.Filling
	}

	_, updateErr := conn.tank.UpdateTankWaterLevel(ctx, connection, input)

	if updateErr.HasError() {
		err.AppendUsecaseError(ErrWaterTankErrorServerError(updateErr.EntityError().Error()))
		return
	}

	return
}
