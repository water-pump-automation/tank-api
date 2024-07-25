package get_group

import (
	"context"
	"time"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type GetGroupWaterTank struct {
	tank water_tank.IWaterTankDatabase
}

func NewGetGroupWaterTank(tank water_tank.IWaterTankDatabase) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (response *ports.WaterTankGroupState, err stack.Error) {
	var states []*water_tank.WaterTank
	response = new(ports.WaterTankGroupState)

	if input.Group != "" {
		states, err = conn.tank.GetTankGroupState(ctx, connection, input)
	} else {
		err.AppendUsecaseError(ErrWaterTankMissingGroup)
		return
	}

	if err.HasError() {
		err.AppendUsecaseError(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if len(states) == 0 {
		err.AppendUsecaseError(ErrWaterTankErrorGroupNotFound(input.Group))
		return
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
