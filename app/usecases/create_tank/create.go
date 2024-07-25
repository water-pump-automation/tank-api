package create_tank

import (
	"context"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/ports"
)

type CreateWaterTank struct {
	tank       water_tank.IWaterTankDatabase
	getUsecase ports.IGetCapacity
}

func NewWaterTank(tank water_tank.IWaterTankDatabase, getUsecase ports.IGetCapacity) *CreateWaterTank {
	return &CreateWaterTank{
		tank:       tank,
		getUsecase: getUsecase,
	}
}

func (conn *CreateWaterTank) Create(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (response *ports.WaterTankState, err error) {
	response = new(ports.WaterTankState)

	if validationErr, err := validation.Validate(ctx, input, validation.CreateTankSchemaLoader); err != nil {
		return nil, err
	} else if validationErr != nil {
		return nil, validationErr
	}

	_, err = conn.getUsecase.GetMaximumCapacity(ctx, connection, &water_tank.GetWaterTankState{
		TankName: input.TankName,
		Group:    input.Group,
	})
	if err != nil {
		return nil, ErrWaterTankAlreadyExists
	}

	if input.MaximumCapacity <= 0 {
		return nil, ErrWaterTankMaximumCapacityZero
	}

	if input.TankName == "" {
		return nil, ErrWaterTankInvalidName
	}

	if input.Group == "" {
		return nil, ErrWaterTankInvalidGroup
	}

	tankSate, createErr := conn.tank.CreateWaterTank(ctx, connection, input)

	if createErr != nil {
		return nil, ErrWaterTankErrorServerError(createErr.Error())
	}

	response.Name = tankSate.Name
	response.Group = tankSate.Group
	response.MaximumCapacity = ports.ConvertCapacityToLiters(tankSate.MaximumCapacity)
	response.TankState = ports.MapTankStateEnum(tankSate.TankState)
	response.CurrentWaterLevel = ports.ConvertCapacityToLiters(tankSate.CurrentWaterLevel)
	response.LastFullTime = tankSate.LastFullTime

	return
}
