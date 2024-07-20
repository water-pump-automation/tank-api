package get_tank

import (
	"time"
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
)

type GetWaterTank struct {
	tank water_tank.WaterTankData
}

func NewGetWaterTank(tank water_tank.WaterTankData) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) GetData(tank string, group string) (MaximumCapacity water_tank.Capacity, accessToken access.AccessToken, err stack.ErrorStack) {
	var state *water_tank.WaterTank
	state, err = conn.tank.GetWaterTankState(group, tank)

	if err.HasError() {
		err.Append(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.Append(ErrWaterTankErrorNotFound(tank))
		return
	}

	return state.MaximumCapacity, state.Access, err
}

func (conn *GetWaterTank) Get(name string, group string) (response *water_tank.WaterTankState, err stack.ErrorStack) {
	response = new(water_tank.WaterTankState)
	var state *water_tank.WaterTank

	state, err = conn.tank.GetWaterTankState(group, name)

	if err.HasError() {
		err.Append(ErrWaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.Append(ErrWaterTankErrorNotFound(name))
		return
	}

	response.Name = state.Name
	response.Group = state.Group
	response.MaximumCapacity = water_tank.ConvertCapacityToLiters(state.MaximumCapacity)
	response.TankState = water_tank.MapTankStateEnum(state.TankState)
	response.CurrentWaterLevel = water_tank.ConvertCapacityToLiters(state.CurrentWaterLevel)
	response.LastFullTime = state.LastFullTime

	now := time.Now()
	response.Datetime = &now

	return
}