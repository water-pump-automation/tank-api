package data

type WaterTankState struct {
	// Fill
}

type WaterTankData interface {
	GetDataByName(names ...string) (state *WaterTankState, err error)
	GetDataByGroup(groups ...string) (state []*WaterTankState, err error)
}
