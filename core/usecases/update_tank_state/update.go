package update_tank_state

import (
	"time"
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
	data "water-tank-api/core/entity/water_tank"
	get_tank "water-tank-api/core/usecases/get/tank"
	"water-tank-api/core/usecases/tank"
)

type UpdateWaterTank struct {
	tank     data.WaterTankData
	capacity tank.Tank
}

func NewWaterTankUpdate(tank data.WaterTankData) *UpdateWaterTank {
	return &UpdateWaterTank{
		tank:     tank,
		capacity: get_tank.NewGetWaterTank(tank),
	}
}

func (conn *UpdateWaterTank) Update(tank string, currentLevel data.Capacity) (err stack.ErrorStack) {
	var maximumCapacity water_tank.Capacity
	var fillState water_tank.State

	maximumCapacity, err = conn.capacity.GetCapacity(tank)

	if err.HasError() {
		if entity := err.EntityError(); entity != nil {
			err.Append(WaterTankErrorServerError(entity.Error()))
			return
		}

		err.Append(WaterTankErrorNotFound(err.LastError().Error()))
		return
	}

	if currentLevel < 0 {
		err.Append(WaterTankCurrentWaterLevelSmallerThanZero)
		return
	}

	if currentLevel > maximumCapacity {
		err.Append(WaterTankCurrentWaterLevelBiggerThanMax)
		return
	} else if currentLevel == maximumCapacity {
		fillState = data.Full
	} else if currentLevel < maximumCapacity {
		fillState = data.Filling
	}

	_, updateErr := conn.tank.UpdateWaterTankState(tank, currentLevel, fillState)

	if updateErr.HasError() {
		err.Append(WaterTankErrorServerError(updateErr.EntityError().Error()))
	}

	if fillState == data.Full {
		_, notifyErr := conn.tank.NotifyFullTank(tank, time.Now())
		if notifyErr.HasError() {
			err.Append(WaterTankErrorServerError(updateErr.EntityError().Error()))
		}
	}

	return
}
