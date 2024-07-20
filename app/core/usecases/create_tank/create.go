package create_tank

import (
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/get_tank"
)

type WaterTank struct {
	tank       water_tank.WaterTankData
	getUsecase GetUsecase
}

func NewWaterTank(tank water_tank.WaterTankData) *WaterTank {
	return &WaterTank{
		tank:       tank,
		getUsecase: get_tank.NewGetWaterTank(tank),
	}
}

func (conn *WaterTank) Create(tank string, group string, capacity water_tank.Capacity) (accessToken access.AccessToken, err stack.ErrorStack) {
	_, _, err = conn.getUsecase.GetData(tank, group)

	if !err.HasError() {
		err.Append(ErrWaterTankAlreadyExists)
		return
	}

	err.PopError()

	if capacity <= 0 {
		err.Append(ErrWaterTankMaximumCapacityZero)
		return
	}

	if tank == "" {
		err.Append(ErrWaterTankInvalidName)
		return
	}

	if group == "" {
		err.Append(ErrWaterTankInvalidGroup)
		return
	}

	accessToken = access.GenerateAccessToken()

	createErr := conn.tank.CreateWaterTank(tank, group, accessToken, capacity)

	if createErr.HasError() {
		err.Append(ErrWaterTankErrorServerError(createErr.EntityError().Error()))
	}

	return
}
