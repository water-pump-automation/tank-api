package update_tank_state

import (
	"context"
	"fmt"
	"tank-api/app/entity/logs"
	"tank-api/app/entity/tank"
	"tank-api/app/usecases/ports"
	"tank-api/app/usecases/validate"
)

type UpdateTank struct {
	tank       tank.ITankDatabase
	getUsecase ports.IGetCapacity
}

func NewTankUpdate(tank tank.ITankDatabase, getUsecase ports.IGetCapacity) *UpdateTank {
	return &UpdateTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *UpdateTank) Update(ctx context.Context, input ports.UsecaseInput) (err error) {
	var maximumCapacity tank.Capacity
	var databaseInput tank.UpdateLevelInput

	if err := validate.ValidateInput(input, &databaseInput, UpdateTankSchemaLoader); err != nil {
		return err
	}

	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s',  level to %s",
			databaseInput.TankName, databaseInput.Group, ports.ConvertCapacityToLiters(databaseInput.NewLevel)),
	)

	maximumCapacity, err = conn.getUsecase.GetMaximumCapacity(ctx, &tank.GetTankStateInput{
		TankName: databaseInput.TankName,
		Group:    databaseInput.Group,
	})

	if err != nil {
		if maximumCapacity == ports.INVALID_CAPACITY {
			return ErrTankErrorServerError(err.Error())
		}

		return ErrTankErrorNotFound
	}

	if databaseInput.NewLevel < 0 {
		return ErrTankCurrentLevelSmallerThanZero
	}

	if ports.MapState(databaseInput.NewLevel, maximumCapacity) == tank.Invalid {
		return ErrTankCurrentLevelBiggerThanMax
	}

	_, updateErr := conn.tank.UpdateTankLevel(ctx, &databaseInput)

	if updateErr != nil {
		return ErrTankErrorServerError(updateErr.Error())
	}

	return
}
