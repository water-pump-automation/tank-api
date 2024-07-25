package create_tank

import (
	"context"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type CreateWaterTank struct {
	tank       water_tank.IWaterTankDatabase
	getUsecase ports.IGetCapacity
}

func NewWaterTank(tank water_tank.IWaterTankDatabase, getUsecase ports.IGetCapacity) *CreateWaterTank {
	return &CreateWaterTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *CreateWaterTank) Create(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (response *ports.WaterTankState, err stack.Error) {
	response = new(ports.WaterTankState)
	_, err = conn.getUsecase.GetMaximumCapacity(ctx, connection, &water_tank.GetWaterTankState{
		TankName: input.TankName,
		Group:    input.Group,
	})

	if !err.HasError() {
		err.AppendUsecaseError(ErrWaterTankAlreadyExists)
		return
	}
	err.PopUsecaseError()

	if input.MaximumCapacity <= 0 {
		err.AppendUsecaseError(ErrWaterTankMaximumCapacityZero)
		return
	}

	if input.TankName == "" {
		err.AppendUsecaseError(ErrWaterTankInvalidName)
		return
	}

	if input.Group == "" {
		err.AppendUsecaseError(ErrWaterTankInvalidGroup)
		return
	}

	tankSate, createErr := conn.tank.CreateWaterTank(ctx, connection, input)

	if createErr.HasError() {
		err.AppendUsecaseError(ErrWaterTankErrorServerError(createErr.EntityError().Error()))
	}

	response.Name = tankSate.Name
	response.Group = tankSate.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(tankSate.MaximumCapacity)
	response.TankState = ports.MapTankStateEnum(tankSate.TankState)
	response.CurrentWaterLevel = ports.ConvertCapacityToLiters(tankSate.CurrentWaterLevel)
	response.LastFullTime = tankSate.LastFullTime

	return
}
