package create_tank

import (
	"water-tank-api/core/entity/data"
	"water-tank-api/core/usecases"
	get_tank "water-tank-api/core/usecases/get"
)

type WaterTank struct {
	tank       data.WaterTankData
	getUsecase *get_tank.GetWaterTank
}

func NewWaterTank(tank data.WaterTankData) *WaterTank {
	return &WaterTank{
		tank:       tank,
		getUsecase: get_tank.NewGetWaterTank(tank),
	}
}

func (conn *WaterTank) Create(tank string, group string, capacity data.Capacity) (errStack usecases.UsecaseErrorStack) {
	tankState, _ := conn.getUsecase.Get(tank)

	if tankState != nil {
		errStack.Append(nil)
		errStack.Append(WaterTankAlreadyExists)
		return
	}

	if capacity < 0 {
		errStack.Append(nil)
		errStack.Append(WaterTankMaximumCapacitySmallerThanZero)
		return
	}

	if tank == "" {
		errStack.Append(nil)
		errStack.Append(WaterTankInvalidName)
		return
	}

	if group == "" {
		errStack.Append(nil)
		errStack.Append(WaterTankInvalidGroup)
		return
	}

	createErr := conn.tank.CreateWaterTank(tank, group, capacity)

	if createErr != nil {
		errStack.Append(createErr)
		errStack.Append(WaterTankErrorServerError(createErr.Error()))
	}

	return
}
