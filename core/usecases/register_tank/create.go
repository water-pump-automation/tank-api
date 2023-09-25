package register_tank

import (
	"water-tank-api/core/entity/access"
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
	get_tank "water-tank-api/core/usecases/get/tank"
	tank "water-tank-api/core/usecases/get_data_interface"
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

func (conn *WaterTank) Create(tank string, group string, capacity data.Capacity) (accessToken access.AccessToken, err stack.ErrorStack) {
	_, _, err = conn.getUsecase.GetData(tank, group)

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

	accessToken = access.GenerateAccessToken()

	createErr := conn.tank.CreateWaterTank(tank, group, accessToken, capacity)

	if createErr.HasError() {
		err.Append(WaterTankErrorServerError(createErr.EntityError().Error()))
	}

	return
}
