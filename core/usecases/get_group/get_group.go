package get_group

import (
	"time"
	"water-tank-api/core/entity/data"
	"water-tank-api/core/usecases"
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

func (conn *GetGroupWaterTank) Get(name string) (response *WaterTankGroupState, err usecases.UsecaseErrorStack) {
	if name == "" {
		name = ALL_GROUPS
	}

	states, entityErr := conn.tank.GetDataByGroup(name)

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
