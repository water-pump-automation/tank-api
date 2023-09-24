package get_group

import (
	"time"
	stack "water-tank-api/core/entity/error_stack"
	"water-tank-api/core/entity/water_tank"
	data "water-tank-api/core/entity/water_tank"
)

const ALL_GROUPS = "ALL"

type GetGroupWaterTank struct {
	tank data.WaterTankData
}

func NewGetGroupWaterTank(tank data.WaterTankData) *GetGroupWaterTank {
	return &GetGroupWaterTank{
		tank: tank,
	}
}

func (conn *GetGroupWaterTank) Get(name string) (response *WaterTankGroupState, err stack.ErrorStack) {
	var state []*water_tank.WaterTankState

	if name == "" {
		name = ALL_GROUPS
	}

	state, err = conn.tank.GetTankGroupState(name)

	if err.HasError() {
		err.Append(WaterTankErrorServerError(err.EntityError().Error()))
		return
	}

	if len(state) == 0 {
		err.Append(WaterTankErrorGroupNotFound(name))
		return
	}

	response.Tank = state
	response.Datetime = time.Now()
	return
}
