package get_group

import (
	"time"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/ports"
)

type GetGroupWaterTank struct {
	tank water_tank.WaterTankData
}

func NewGetGroupWaterTank(tank water_tank.WaterTankData) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(name string) (response *ports.WaterTankGroupState, err stack.ErrorStack) {
	var states []*water_tank.WaterTank
	response = new(ports.WaterTankGroupState)

	if name != "" {
		states, err = conn.tank.GetTankGroupState(name)
	} else {
		err.Append(ErrWaterTankMissingGroup)
		return
	}

	if err.HasError() {
		err.Append(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if len(states) == 0 {
		err.Append(ErrWaterTankErrorGroupNotFound(name))
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
