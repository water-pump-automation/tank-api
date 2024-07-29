package get_tank

import (
	"context"
	"fmt"
	"time"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
	"water-tank-api/app/usecases/validate"
)

type GetWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetWaterTank(tank water_tank.IWaterTankDatabase) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) Exists(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (bool, error) {
	state, err := conn.tank.GetWaterTankState(ctx, connection, input)
	if err != nil {
		return false, ErrWaterTankErrorServerError(err.Error())
	}

	if state != nil {
		return true, nil
	}
	return false, nil
}

func (conn *GetWaterTank) GetMaximumCapacity(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankStateInput) (maximumCapacity water_tank.Capacity, err error) {
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

func (conn *GetWaterTank) Get(ctx context.Context, connection water_tank.IConn, input ports.UsecaseInput) (response *ports.WaterTankState, err error) {
	response = new(ports.WaterTankState)
	var state *water_tank.WaterTank
	var databaseInput water_tank.GetWaterTankStateInput

	if err := validate.ValidateInput(ctx, input, &databaseInput, GetTankSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank state, of group '%s'...", databaseInput.TankName, databaseInput.Group))

	state, err = conn.tank.GetWaterTankState(ctx, connection, &databaseInput)

	if err != nil {
		return nil, ErrWaterTankErrorServerError(err.Error())
	}

	if state == nil {
		return nil, ErrWaterTankErrorNotFound
	}

	response.Name = state.Name
	response.Group = state.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(state.MaximumCapacity)
	response.TankState = ports.ConvertState(ports.MapWaterState(state.CurrentWaterLevel, state.MaximumCapacity))
	response.CurrentWaterLevel = ports.ConvertCapacityToLiters(state.CurrentWaterLevel)
	response.LastFullTime = state.LastFullTime

	now := time.Now()
	response.Datetime = &now

	return
}
