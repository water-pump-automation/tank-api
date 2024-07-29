package create_tank

import (
	"context"
	"fmt"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
	"water-tank-api/app/usecases/validate"
)

type CreateWaterTank struct {
	tank       water_tank.IWaterTankDatabase
	tankExists ports.ITankExists
}

func NewWaterTank(tank water_tank.IWaterTankDatabase, tankExists ports.ITankExists) *CreateWaterTank {
	return &CreateWaterTank{
		tank:       tank,
		tankExists: tankExists,
	}
}

func (conn *CreateWaterTank) Create(ctx context.Context, connection water_tank.IConn, input ports.UsecaseInput) (response *ports.WaterTankState, err error) {
	response = new(ports.WaterTankState)
	var databaseInput water_tank.CreateInput

	if err := validate.ValidateInput(ctx, input, &databaseInput, CreateTankSchemaLoader); err != nil {
		return nil, err
	}

	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			databaseInput.TankName, databaseInput.Group, ports.ConvertCapacityToLiters(databaseInput.MaximumCapacity)),
	)

	exists, err := conn.tankExists.Exists(ctx, connection, &water_tank.GetWaterTankStateInput{
		TankName: databaseInput.TankName,
		Group:    databaseInput.Group,
	})
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrWaterTankAlreadyExists
	}

	if databaseInput.MaximumCapacity <= 0 {
		return nil, ErrWaterTankMaximumCapacityZero
	}

	if databaseInput.TankName == "" {
		return nil, ErrWaterTankInvalidName
	}

	if databaseInput.Group == "" {
		return nil, ErrWaterTankInvalidGroup
	}

	tankSate, createErr := conn.tank.CreateWaterTank(ctx, connection, &databaseInput)

	if createErr != nil {
		return nil, ErrWaterTankErrorServerError(createErr.Error())
	}

	response.Name = tankSate.Name
	response.Group = tankSate.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(tankSate.MaximumCapacity)
	response.TankState = ports.EMPTY
	response.CurrentWaterLevel = ports.ConvertCapacityToLiters(tankSate.CurrentWaterLevel)
	response.LastFullTime = tankSate.LastFullTime

	return
}
