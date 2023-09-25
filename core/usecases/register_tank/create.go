package register_tank

import (
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
	get_tank "water-tank-api/core/usecases/get/tank"
	"water-tank-api/core/usecases/tank"
)

type WaterTank struct {
	tank       data.WaterTankData
	getUsecase tank.Tank
}

func NewWaterTank(tank data.WaterTankData) *WaterTank {
	return &WaterTank{
		tank:       tank,
		getUsecase: get_tank.NewGetWaterTank(tank),
	}
}

func (conn *WaterTank) Create(tank string, group string, capacity data.Capacity) (err stack.ErrorStack) {
	_, err = conn.getUsecase.GetCapacity(tank)

	if !err.HasError() {
		err.Append(WaterTankAlreadyExists)
		return
	}

	err.PopError()

	if capacity <= 0 {
		err.Append(WaterTankMaximumCapacityZero)
		return
	}

	if tank == "" {
		err.Append(WaterTankInvalidName)
		return
	}

	if group == "" {
		err.Append(WaterTankInvalidGroup)
		return
	}

	createErr := conn.tank.CreateWaterTank(tank, group, capacity)

	if createErr.HasError() {
		err.Append(WaterTankErrorServerError(createErr.EntityError().Error()))
	}

	return
}
