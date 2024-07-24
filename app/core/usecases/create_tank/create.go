package create_tank

import (
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type CreateWaterTank struct {
	tank       water_tank.WaterTankData
	getUsecase ports.GetUsecase
}

func NewWaterTank(tank water_tank.WaterTankData, getUsecase ports.GetUsecase) *CreateWaterTank {
	return &CreateWaterTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *CreateWaterTank) Create(tank string, group string, capacity water_tank.Capacity) (accessToken access.AccessToken, err stack.ErrorStack) {
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
