package tank

import (
	"time"
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

func (conn *GetWaterTank) GetCapacity(tank string) (MaximumCapacity water_tank.Capacity, err stack.ErrorStack) {
	state, err := conn.Get(tank)

	return get.ConverLitersToCapacity(state.MaximumCapacity), err
}

func (conn *GetWaterTank) Get(name string) (response *get.WaterTankState, err stack.ErrorStack) {
	response = new(get.WaterTankState)
	var state *water_tank.WaterTank

	state, err = conn.tank.GetWaterTankState(name)

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
