package water_tank

import (
	"context"
	stack "water-tank-api/app/core/entity/error_stack"
)

type CreateInput struct {
	TankName        string   `json:"tank_name"`
	Group           string   `json:"group"`
	MaximumCapacity Capacity `json:"maximum_capacity"`
}

type UpdateWaterLevelInput struct {
	TankName      string
	Group         string
	State         State
	NewWaterLevel Capacity `json:"water_level"`
}

type GetWaterTankState struct {
	Group    string
	TankName string
}

type GetGroupTanks struct {
	Group string
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
	CreateWaterTank(ctx context.Context, connection IConn, input *CreateInput) (tank *WaterTank, err stack.Error)
	UpdateTankWaterLevel(ctx context.Context, connection IConn, input *UpdateWaterLevelInput) (tank *WaterTank, err stack.Error)
	GetWaterTankState(ctx context.Context, connection IConn, input *GetWaterTankState) (tank *WaterTank, err stack.Error)
	GetTankGroupState(ctx context.Context, connection IConn, input *GetGroupTanks) (tank []*WaterTank, err stack.Error)
}
