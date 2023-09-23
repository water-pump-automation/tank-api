package data

type State int

const (
	Empty State = iota
	Filling
	Full
)

type Capacity float32

type WaterTankState struct {
	Name              string
	MaximumCapacity   Capacity
	TankState         State
	CurrentWaterLevel Capacity
	Water             *Water
}

type WaterTankData interface {
	CreateWaterTank(name string, group string) (err error)
	UpdateWaterTankState(name string, waterLevel *Capacity) (state *WaterTankState, err error)
	GetDataByName(names ...string) (state *WaterTankState, err error)
	GetDataByGroup(groups ...string) (state []*WaterTankState, err error)
}
