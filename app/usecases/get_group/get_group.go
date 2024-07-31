package get_group

import (
	"context"
	"fmt"
	"tank-api/app/entity/logs"
	"tank-api/app/entity/tank"
	"tank-api/app/usecases/ports"
	"tank-api/app/usecases/validate"
	"time"
)

type GetGroupTank struct {
	tank tank.ITankDatabase
}

func NewGetGroupTank(tank tank.ITankDatabase) *GetGroupTank {
	return &GetGroupTank{
		tank: tank,
	}
}

func (conn *GetGroupTank) Get(ctx context.Context, input ports.UsecaseInput) (response *ports.TankGroupState, err error) {
	var databaseInput tank.GetGroupTanksInput
	var states []*tank.Tank
	response = new(ports.TankGroupState)

	if err := validate.ValidateInput(input, &databaseInput, GetGroupSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank group...", databaseInput.Group))

	states, err = conn.tank.GetTankGroupState(ctx, &databaseInput)
	if err != nil {
		return nil, ErrTankErrorServerError(err.Error())
	}

	if len(states) == 0 {
		return nil, ErrTankErrorGroupNotFound
	}

	for _, tank := range states {
		state := new(ports.TankState)
		state.Name = tank.Name
		state.Group = tank.Group
		state.MaximumCapacity = ports.ConvertCapacityToLiters(tank.MaximumCapacity)
		state.TankState = ports.ConvertState(ports.MapState(tank.CurrentLevel, tank.MaximumCapacity))
		state.CurrentLevel = ports.ConvertCapacityToLiters(tank.CurrentLevel)
		state.LastFullTime = tank.LastFullTime

		response.Tanks = append(response.Tanks, state)
	}
	response.Datetime = time.Now()
	return
}
