package update_tank_state

import (
	"context"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
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

func (conn *UpdateWaterTank) Update(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (err error) {
	var maximumCapacity water_tank.Capacity

	if validationErr, err := validation.Validate(ctx, input, validation.UpdateTankSchemaLoader); err != nil {
		return err
	} else if validationErr != nil {
		return validationErr
	}

	maximumCapacity, err = conn.getUsecase.GetMaximumCapacity(ctx, connection, &water_tank.GetWaterTankState{
		TankName: input.TankName,
		Group:    input.Group,
	})

	if err != nil {
		if maximumCapacity == ports.INVALID_CAPACITY {
			return ErrWaterTankErrorServerError(err.Error())
		}

		return ErrWaterTankErrorNotFound
	}

	if input.NewWaterLevel < 0 {
		return ErrWaterTankCurrentWaterLevelSmallerThanZero
	}

	if input.NewWaterLevel > maximumCapacity {
		return ErrWaterTankCurrentWaterLevelBiggerThanMax
	} else if input.NewWaterLevel == maximumCapacity {
		input.State = water_tank.Full
	} else if input.NewWaterLevel == 0 {
		input.State = water_tank.Empty
	} else if input.NewWaterLevel < maximumCapacity {
		input.State = water_tank.Filling
	}

	_, updateErr := conn.tank.UpdateTankWaterLevel(ctx, connection, input)

	if updateErr != nil {
		return ErrWaterTankErrorServerError(updateErr.Error())
	}

	return
}
