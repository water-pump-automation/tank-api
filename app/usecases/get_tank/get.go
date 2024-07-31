package get_tank

import (
	"context"
	"fmt"
	"tank-api/app/entity/logs"
	"tank-api/app/entity/tank"
	"tank-api/app/usecases/ports"
	"tank-api/app/usecases/validate"
	"time"
)

type GetTank struct {
	tank tank.ITankDatabase
}

func NewGetTank(tank tank.ITankDatabase) *GetTank {
	return &GetTank{
		tank: tank,
	}
}

func (conn *GetTank) Exists(ctx context.Context, input *tank.GetTankStateInput) (bool, error) {
	state, err := conn.tank.GetTankState(ctx, input)
	if err != nil {
		return false, ErrTankErrorServerError(err.Error())
	}

	if state != nil {
		return true, nil
	}
	return false, nil
}

func (conn *GetTank) GetMaximumCapacity(ctx context.Context, input *tank.GetTankStateInput) (maximumCapacity tank.Capacity, err error) {
	var state *tank.Tank
	state, err = conn.tank.GetTankState(ctx, input)

	if err != nil {
		return ports.INVALID_CAPACITY, ErrTankErrorServerError(err.Error())
	}

	if state == nil {
		return ports.EMPTY_CAPACITY, ErrTankErrorNotFound
	}

	return state.MaximumCapacity, err
}

func (conn *GetTank) Get(ctx context.Context, input ports.UsecaseInput) (response *ports.TankState, err error) {
	response = new(ports.TankState)
	var state *tank.Tank
	var databaseInput tank.GetTankStateInput

	if err := validate.ValidateInput(ctx, input, &databaseInput, GetTankSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank state, of group '%s'...", databaseInput.TankName, databaseInput.Group))

	state, err = conn.tank.GetTankState(ctx, &databaseInput)

	if err != nil {
		return nil, ErrTankErrorServerError(err.Error())
	}

	if state == nil {
		return nil, ErrTankErrorNotFound
	}

	response.Name = state.Name
	response.Group = state.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(state.MaximumCapacity)
	response.TankState = ports.ConvertState(ports.MapState(state.CurrentLevel, state.MaximumCapacity))
	response.CurrentLevel = ports.ConvertCapacityToLiters(state.CurrentLevel)
	response.LastFullTime = state.LastFullTime
	response.Datetime = time.Now()

	return
}
