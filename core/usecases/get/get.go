package get_tank

import (
	"time"
	"water-tank-api/core/entity/data"
	"water-tank-api/core/usecases"
)

type GetWaterTank struct {
	tank data.WaterTankData
}

func NewGetWaterTank(tank data.WaterTankData) *GetWaterTank {
	return &GetWaterTank{
		tank: tank,
	}
}

func (conn *GetWaterTank) Get(name string) (response *WaterTankState, err usecases.UsecaseErrorStack) {
	state, entityErr := conn.tank.GetWaterTankState(name)

	if entityErr != nil {
		err.Append(entityErr)
		err.Append(WaterTankErrorServerError(entityErr.Error()))
		return
	}

	if state == nil {
		err.Append(nil)
		err.Append(WaterTankErrorNotFound(name))
		return
	}

	response.Tank = state
	response.Datetime = time.Now()

	return
}
