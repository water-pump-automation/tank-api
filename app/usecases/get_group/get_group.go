package get_group

import (
	"context"
	"fmt"
	"time"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
	"water-tank-api/app/usecases/validate"
)

type GetGroupWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetGroupWaterTank(tank water_tank.IWaterTankDatabase) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(ctx context.Context, connection water_tank.IConn, input ports.UsecaseInput) (response *ports.WaterTankGroupState, err error) {
	var databaseInput water_tank.GetGroupTanksInput
	var states []*water_tank.WaterTank
	response = new(ports.WaterTankGroupState)

	if err := validate.ValidateInput(ctx, input, &databaseInput, GetGroupSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank group...", databaseInput.Group))

	states, err = conn.tank.GetTankGroupState(ctx, connection, &databaseInput)
	if err != nil {
		return nil, ErrWaterTankErrorServerError(err.Error())
	}

	if len(states) == 0 {
		return nil, ErrWaterTankErrorGroupNotFound
	}

	for _, tank := range states {
		state := new(ports.WaterTankState)
		state.Name = tank.Name
		state.Group = tank.Group
		state.MaximumCapacity = ports.ConvertCapacityToLiters(tank.MaximumCapacity)
		state.TankState = ports.ConvertState(ports.MapWaterState(tank.CurrentWaterLevel, tank.MaximumCapacity))
		state.CurrentWaterLevel = ports.ConvertCapacityToLiters(tank.CurrentWaterLevel)
		state.LastFullTime = tank.LastFullTime

		response.Tanks = append(response.Tanks, state)
	}
	response.Datetime = time.Now()
	return
}
