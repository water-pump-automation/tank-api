package update_tank_state

import (
	"context"
	"fmt"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
	"water-tank-api/app/usecases/validate"
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

func (conn *UpdateWaterTank) Update(ctx context.Context, connection water_tank.IConn, input ports.UsecaseInput) (err error) {
	var maximumCapacity water_tank.Capacity
	var databaseInput water_tank.UpdateWaterLevelInput

	if err := validate.ValidateInput(ctx, input, &databaseInput, UpdateTankSchemaLoader); err != nil {
		return err
	}

	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s', water level to %s",
			databaseInput.TankName, databaseInput.Group, ports.ConvertCapacityToLiters(databaseInput.NewWaterLevel)),
	)

	maximumCapacity, err = conn.getUsecase.GetMaximumCapacity(ctx, connection, &water_tank.GetWaterTankStateInput{
		TankName: databaseInput.TankName,
		Group:    databaseInput.Group,
	})

	if err != nil {
		if maximumCapacity == ports.INVALID_CAPACITY {
			return ErrWaterTankErrorServerError(err.Error())
		}

		return ErrWaterTankErrorNotFound
	}

	if databaseInput.NewWaterLevel < 0 {
		return ErrWaterTankCurrentWaterLevelSmallerThanZero
	}

	databaseInput.State = ports.MapWaterState(databaseInput.NewWaterLevel, maximumCapacity)
	if databaseInput.State == water_tank.Invalid {
		return ErrWaterTankCurrentWaterLevelBiggerThanMax
	}

	_, updateErr := conn.tank.UpdateTankWaterLevel(ctx, connection, &databaseInput)

	if updateErr != nil {
		return ErrWaterTankErrorServerError(updateErr.Error())
	}

	return
}
