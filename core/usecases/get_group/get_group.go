package get_group

import (
	"time"
	stack "water-tank-api/core/entity/error_stack"
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
	if name == "" {
		name = ALL_GROUPS
	}

	states, entityErr := conn.tank.GetTankGroupState(name)

	if entityErr != nil {
		err.Append(entityErr)
		err.Append(WaterTankErrorServerError(entityErr.Error()))
		return
	}

	if len(states) == 0 {
		err.Append(nil)
		err.Append(WaterTankErrorGroupNotFound(name))
		return
	}

	response.Tank = states
	response.Datetime = time.Now()
	return
}
