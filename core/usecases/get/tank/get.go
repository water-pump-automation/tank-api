package tank

import (
	"time"
	"water-tank-api/core/entity/access"
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
	data "water-tank-api/core/entity/water_tank"
	"water-tank-api/core/usecases/get"
)

type GetWaterTank struct {
	tank data.WaterTankData
}

func NewGetWaterTank(tank data.WaterTankData) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) GetData(tank string, group string) (MaximumCapacity water_tank.Capacity, accessToken access.AccessToken, err stack.ErrorStack) {
	var state *water_tank.WaterTank
	state, err = conn.tank.GetWaterTankState(group, tank)

	if err.HasError() {
		err.Append(WaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.Append(WaterTankErrorNotFound(tank))
		return
	}

	return state.MaximumCapacity, state.Access, err
}

func (conn *GetWaterTank) Get(name string, group string) (response *get.WaterTankState, err stack.ErrorStack) {
	response = new(get.WaterTankState)
	var state *water_tank.WaterTank

	state, err = conn.tank.GetWaterTankState(group, name)

	if err.HasError() {
		err.Append(WaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.Append(WaterTankErrorNotFound(name))
		return
	}

	response.Name = state.Name
	response.Group = state.Group
	response.MaximumCapacity = get.ConvertCapacityToLiters(state.MaximumCapacity)
	response.TankState = get.MapTankStateEnum(state.TankState)
	response.CurrentWaterLevel = get.ConvertCapacityToLiters(state.CurrentWaterLevel)
	response.LastFullTime = state.LastFullTime

	now := time.Now()
	response.Datetime = &now

	return
}
