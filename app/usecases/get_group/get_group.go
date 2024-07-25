package get_group

import (
	"context"
	"time"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
)

type GetGroupWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetGroupWaterTank(tank water_tank.IWaterTankDatabase) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (response *ports.WaterTankGroupState, err error) {
	var states []*water_tank.WaterTank
	response = new(ports.WaterTankGroupState)

	if validationErr, err := validation.Validate(ctx, input, validation.GetGroupSchemaLoader); err != nil {
		return nil, err
	} else if validationErr != nil {
		return nil, validationErr
	}

	states, err = conn.tank.GetTankGroupState(ctx, connection, input)
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
		state.TankState = ports.MapTankStateEnum(tank.TankState)
		state.CurrentWaterLevel = ports.ConvertCapacityToLiters(tank.CurrentWaterLevel)
		state.LastFullTime = tank.LastFullTime

		response.Tanks = append(response.Tanks, state)
	}
	response.Datetime = time.Now()
	return
}
