package get_tank

import (
	"time"
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
	data "water-tank-api/core/entity/water_tank"
)

type GetWaterTank struct {
	tank data.WaterTankData
}

func NewGetWaterTank(tank data.WaterTankData) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) Get(name string) (response *WaterTankState, err stack.ErrorStack) {
	var state *water_tank.WaterTankState

	state, err = conn.tank.GetWaterTankState(name)

	if err.HasError() {
		err.Append(WaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if state == nil {
		err.Append(WaterTankErrorNotFound(name))
		return
	}

	response.Tank = state
	response.Datetime = time.Now()

	return
}
