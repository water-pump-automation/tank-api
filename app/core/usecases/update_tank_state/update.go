package update_tank_state

import (
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type UpdateWaterTank struct {
	tank       water_tank.WaterTankData
	getUsecase ports.GetUsecase
}

func NewWaterTankUpdate(tank water_tank.WaterTankData, getUsecase ports.GetUsecase) *UpdateWaterTank {
	return &UpdateWaterTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *UpdateWaterTank) Update(tank string, group string, accessToken access.AccessToken, currentLevel water_tank.Capacity) (err stack.ErrorStack) {
	var maximumCapacity water_tank.Capacity
	var fillState water_tank.State
	var existingAccessToken access.AccessToken

	maximumCapacity, existingAccessToken, err = conn.getUsecase.GetData(tank, group)

	if err.HasError() {
		if entity := err.EntityError(); entity != nil {
			err.Append(ErrWaterTankErrorServerError(entity.Error()))
			return
		}

		err.Append(ErrWaterTankErrorNotFound(tank))
		return
	}

	if currentLevel < 0 {
		err.Append(ErrWaterTankCurrentWaterLevelSmallerThanZero)
		return
	}

	if currentLevel > maximumCapacity {
		err.Append(ErrWaterTankCurrentWaterLevelBiggerThanMax)
		return
	} else if currentLevel == maximumCapacity {
		fillState = water_tank.Full
	} else if currentLevel == 0 {
		fillState = water_tank.Empty
	} else if currentLevel < maximumCapacity {
		fillState = water_tank.Filling
	}

	if accessToken != existingAccessToken {
		err.Append(ErrWaterTankInvalidToken)
		return
	}

	_, updateErr := conn.tank.UpdateWaterTankState(tank, group, currentLevel, fillState)

	if updateErr.HasError() {
		err.Append(ErrWaterTankErrorServerError(updateErr.EntityError().Error()))
		return
	}

	return
}
