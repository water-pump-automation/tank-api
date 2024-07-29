package water_tank

import (
	"context"
)

type CreateInput struct {
	TankName        string   `json:"tank_name,omitempty"`
	Group           string   `json:"group,omitempty"`
	MaximumCapacity Capacity `json:"maximum_capacity,omitempty"`
}

type UpdateWaterLevelInput struct {
	TankName      string `json:"tank_name,omitempty"`
	Group         string `json:"group,omitempty"`
	State         State
	NewWaterLevel Capacity `json:"water_level,omitempty"`
}

type GetWaterTankStateInput struct {
	Group    string `json:"group,omitempty"`
	TankName string `json:"tank_name,omitempty"`
}

type GetGroupTanksInput struct {
	Group string `json:"group,omitempty"`
}

type ConnCallback func(ctx context.Context)

type IPool interface {
	Acquire() (IConn, error)
	AcquireTransaction() (IConn, error)
}

type IConn interface {
	Release() error
	Query(ctx context.Context, callback ConnCallback)
}

type IWaterTankDatabase interface {
	CreateWaterTank(ctx context.Context, connection IConn, input *CreateInput) (tank *WaterTank, err error)
	UpdateTankWaterLevel(ctx context.Context, connection IConn, input *UpdateWaterLevelInput) (tank *WaterTank, err error)
	GetWaterTankState(ctx context.Context, connection IConn, input *GetWaterTankStateInput) (tank *WaterTank, err error)
	GetTankGroupState(ctx context.Context, connection IConn, input *GetGroupTanksInput) (tank []*WaterTank, err error)
}
