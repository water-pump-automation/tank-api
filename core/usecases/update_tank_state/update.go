package update_tank_state

import (
	"water-tank-api/core/entity/data"
	"water-tank-api/core/usecases"
	get_tank "water-tank-api/core/usecases/get_tank"
)

type UpdateWaterTank struct {
	tank       data.WaterTankData
	getUsecase *get_tank.GetWaterTank
}

func NewWaterTankUpdate(tank data.WaterTankData) *UpdateWaterTank {
	return &UpdateWaterTank{
		tank:       tank,
		getUsecase: get_tank.NewGetWaterTank(tank),
	}
}

func (conn *UpdateWaterTank) Update(tank string, currentLevel data.Capacity) (errStack usecases.UsecaseErrorStack, foundErr error) {
	tankState, getErr := conn.getUsecase.Get(tank)

	if getErr.HasError() {
		if errStack.EntityError() != nil {
			errStack.Append(getErr.EntityError())
			errStack.Append(getErr.UsecaseError())
			return
		}

		foundErr = errStack.UsecaseError()
		return
	}

	if currentLevel > tankState.Tank.MaximumCapacity {
		errStack.Append(nil)
		errStack.Append(WaterTankCurrentWaterLevelBiggerThanMax)
		return
	}

	if currentLevel < 0 {
		errStack.Append(nil)
		errStack.Append(WaterTankCurrentWaterLevelSmallerThanZero)
		return
	}

	_, updateErr := conn.tank.UpdateWaterTankState(tank, currentLevel)

	if updateErr != nil {
		errStack.Append(updateErr)
		errStack.Append(WaterTankErrorServerError(updateErr.Error()))
	}

	return
}
