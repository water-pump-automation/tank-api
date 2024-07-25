package get_tank

import (
	"context"
	"time"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
)

type GetWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetWaterTank(tank water_tank.IWaterTankDatabase) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) GetMaximumCapacity(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (maximumCapacity water_tank.Capacity, err error) {
	var state *water_tank.WaterTank
	state, err = conn.tank.GetWaterTankState(ctx, connection, input)

	if err != nil {
		return ports.INVALID_CAPACITY, ErrWaterTankErrorServerError(err.Error())
	}

	if state == nil {
		return ports.EMPTY_CAPACITY, ErrWaterTankErrorNotFound
	}

	return state.MaximumCapacity, err
}

func (conn *GetWaterTank) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (response *ports.WaterTankState, err error) {
	response = new(ports.WaterTankState)
	var state *water_tank.WaterTank

	if validationErr, err := validation.Validate(ctx, input, validation.GetTankSchemaLoader); err != nil {
		return nil, err
	} else if validationErr != nil {
		return nil, validationErr
	}

	state, err = conn.tank.GetWaterTankState(ctx, connection, input)

	if err != nil {
		return nil, ErrWaterTankErrorServerError(err.Error())
	}

	if state == nil {
		return nil, ErrWaterTankErrorNotFound
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
