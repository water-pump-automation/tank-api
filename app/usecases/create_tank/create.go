package create_tank

import (
	"context"
	"fmt"
	"tank-api/app/entity/logs"
	"tank-api/app/entity/tank"
	"tank-api/app/usecases/ports"
	"tank-api/app/usecases/validate"
)

type CreateTank struct {
	tank       tank.ITankDatabase
	tankExists ports.ITankExists
}

func NewTank(tank tank.ITankDatabase, tankExists ports.ITankExists) *CreateTank {
	return &CreateTank{
		tank:       tank,
		tankExists: tankExists,
	}
}

func (conn *CreateTank) Create(ctx context.Context, input ports.UsecaseInput) (response *ports.TankState, err error) {
	response = new(ports.TankState)
	var databaseInput tank.CreateInput

	if err := validate.ValidateInput(ctx, input, &databaseInput, CreateTankSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			databaseInput.TankName, databaseInput.Group, ports.ConvertCapacityToLiters(databaseInput.MaximumCapacity)),
	)

	exists, err := conn.tankExists.Exists(ctx, &tank.GetTankStateInput{
		TankName: databaseInput.TankName,
		Group:    databaseInput.Group,
	})
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrTankAlreadyExists
	}

	if databaseInput.MaximumCapacity <= 0 {
		return nil, ErrTankMaximumCapacityZero
	}

	if databaseInput.TankName == "" {
		return nil, ErrTankInvalidName
	}

	if databaseInput.Group == "" {
		return nil, ErrTankInvalidGroup
	}

	tankState, createErr := conn.tank.CreateTank(ctx, &databaseInput)

	if createErr != nil {
		return nil, ErrTankErrorServerError(createErr.Error())
	}

	response.Name = tankState.Name
	response.Group = tankState.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(tankState.MaximumCapacity)
	response.TankState = ports.ConvertState(ports.MapState(tankState.CurrentLevel, tankState.MaximumCapacity))
	response.CurrentLevel = ports.ConvertCapacityToLiters(tankState.CurrentLevel)
	response.LastFullTime = tankState.LastFullTime

	return
}
