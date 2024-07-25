package get_tank

import (
	"context"
	"time"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type GetWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetWaterTank(tank water_tank.IWaterTankDatabase) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) GetMaximumCapacity(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (maximumCapacity water_tank.Capacity, err stack.Error) {
	var state *water_tank.WaterTank
	state, err = conn.tank.GetWaterTankState(ctx, connection, input)

	if err.HasError() {
		err.AppendUsecaseError(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.AppendUsecaseError(ErrWaterTankErrorNotFound(input.TankName))
		return
	}

	return state.MaximumCapacity, err
}

func (conn *GetWaterTank) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (response *ports.WaterTankState, err stack.Error) {
	response = new(ports.WaterTankState)
	var state *water_tank.WaterTank

	state, err = conn.tank.GetWaterTankState(ctx, connection, input)

	if err.HasError() {
		err.AppendUsecaseError(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.AppendUsecaseError(ErrWaterTankErrorNotFound(input.TankName))
		return
	}

	response.Name = state.Name
	response.Group = state.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(state.MaximumCapacity)
	response.TankState = ports.MapTankStateEnum(state.TankState)
	response.CurrentWaterLevel = ports.ConvertCapacityToLiters(state.CurrentWaterLevel)
	response.LastFullTime = state.LastFullTime

	now := time.Now()
	response.Datetime = &now

	return
}
