package tank

import (
	"context"
)

type CreateInput struct {
	TankName        string   `json:"tank_name,omitempty"`
	Group           string   `json:"group,omitempty"`
	MaximumCapacity Capacity `json:"maximum_capacity,omitempty"`
}

type UpdateLevelInput struct {
	TankName string   `json:"tank_name,omitempty"`
	Group    string   `json:"group,omitempty"`
	NewLevel Capacity `json:"level,omitempty"`
}

type GetTankStateInput struct {
	Group    string `json:"group,omitempty"`
	TankName string `json:"tank_name,omitempty"`
}

type GetGroupTanksInput struct {
	Group string `json:"group,omitempty"`
}

type ITankDatabase interface {
	CreateTank(ctx context.Context, input *CreateInput) (tank *Tank, err error)
	UpdateTankLevel(ctx context.Context, input *UpdateLevelInput) (tank *Tank, err error)
	GetTankState(ctx context.Context, input *GetTankStateInput) (tank *Tank, err error)
	GetTankGroupState(ctx context.Context, input *GetGroupTanksInput) (tank []*Tank, err error)
}
