package postgresql

import "water-tank-api/core/entity/data"

type WaterTankPostgreSQLData struct{}

func (repository *WaterTankPostgreSQLData) GetDataByName(names ...string) (state *data.WaterTankState, err error) {
	return
}
func (repository *WaterTankPostgreSQLData) GetDataByGroup(groups ...string) (state []*data.WaterTankState, err error) {
	return
}
