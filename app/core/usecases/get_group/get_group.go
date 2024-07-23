package get_group

import (
	"time"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases"
)

type GetGroupWaterTank struct {
	tank water_tank.WaterTankData
}

func NewGetGroupWaterTank(tank water_tank.WaterTankData) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(name string) (response *usecases.WaterTankGroupState, err stack.ErrorStack) {
	var states []*water_tank.WaterTank
	response = new(usecases.WaterTankGroupState)

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
		state := new(usecases.WaterTankState)
		state.Name = tank.Name
		state.Group = tank.Group
		state.MaximumCapacity = usecases.ConvertCapacityToLiters(tank.MaximumCapacity)
		state.TankState = usecases.MapTankStateEnum(tank.TankState)
		state.CurrentWaterLevel = usecases.ConvertCapacityToLiters(tank.CurrentWaterLevel)
		state.LastFullTime = tank.LastFullTime

		response.Tanks = append(response.Tanks, state)
	}
	response.Datetime = time.Now()
	return
}
