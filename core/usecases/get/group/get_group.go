package group

import (
	"time"
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/logs"
	"water-tank-api/core/entity/water_tank"
	data "water-tank-api/core/entity/water_tank"
	"water-tank-api/core/usecases/get"
)

type GetGroupWaterTank struct {
	tank data.WaterTankData
}

func NewGetGroupWaterTank(tank data.WaterTankData) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(name string) (response *get.WaterTankGroupState, err stack.ErrorStack) {
	var states []*water_tank.WaterTankState
	response = new(get.WaterTankGroupState)

	if name != "" {
		states, err = conn.tank.GetTankGroupState(name)
	} else {
		logs.Gateway().Info("Empty group! Retrieving all groups...")
		states, err = conn.tank.GetAllTankGroupState()
	}

	if err.HasError() {
		err.Append(WaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if len(states) == 0 {
		err.Append(WaterTankErrorGroupNotFound(name))
		return
	}

	for _, tank := range states {
		state := new(get.WaterTankState)
		state.Name = tank.Name
		state.Group = tank.Group
		state.MaximumCapacity = get.ConvertCapacityToLiters(tank.MaximumCapacity)
		state.TankState = get.MapTankStateEnum(tank.TankState)
		state.CurrentWaterLevel = get.ConvertCapacityToLiters(tank.CurrentWaterLevel)

		response.Tanks = append(response.Tanks, state)
	}
	response.Datetime = time.Now()
	return
}
