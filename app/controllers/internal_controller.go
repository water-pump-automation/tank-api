package controllers

import (
	"context"
	"fmt"
	"water-tank-api/app/controllers/response"
	"water-tank-api/app/controllers/validation"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/create_tank"
	"water-tank-api/app/core/usecases/get_group"
	"water-tank-api/app/core/usecases/get_tank"
	"water-tank-api/app/core/usecases/ports"
	update_tank_state "water-tank-api/app/core/usecases/update_tank_state"
)

type InternalController struct {
	getTankUsecase    *get_tank.GetWaterTank
	getGroupUsecase   *get_group.GetGroupWaterTank
	createTankUsecase *create_tank.CreateWaterTank
	updateTankUsecase *update_tank_state.UpdateWaterTank
}

func NewInternalController(
	getTankUsecase *get_tank.GetWaterTank,
	getGroupUsecase *get_group.GetGroupWaterTank,
	createTankUsecase *create_tank.CreateWaterTank,
	updateTankUsecase *update_tank_state.UpdateWaterTank,
) *InternalController {
	return &InternalController{
		getTankUsecase:    getTankUsecase,
		getGroupUsecase:   getGroupUsecase,
		createTankUsecase: createTankUsecase,
		updateTankUsecase: updateTankUsecase,
	}
}

func (controller *InternalController) Create(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (resp *response.ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			input.TankName, input.Group, ports.ConvertCapacityToLiters(input.MaximumCapacity)),
	)

	resp = validation.Validate(ctx, input, validation.CreateTankSchemaLoader)
	if resp != nil {
		return
	}

	tankState, usecaseErr := controller.createTankUsecase.Create(ctx, connection, input)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			resp = response.NewControllerError(response.WaterTankInvalidRequest, usecaseErr.LastUsecaseError().Error())
		default:
			resp = response.NewControllerError(response.WaterTankInternalServerError, usecaseErr.LastUsecaseError().Error())
		}

		err = usecaseErr.LastUsecaseError()
		logs.Gateway().Error(err.Error())
		return
	}

	resp = response.NewControllerResponse(response.WaterTankOK, tankState)
	return
}

func (controller *InternalController) Update(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (resp *response.ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s', water level to %s",
			input.TankName, input.Group, ports.ConvertCapacityToLiters(input.NewWaterLevel)),
	)

	resp = validation.Validate(ctx, input, validation.UpdateTankSchemaLoader)
	if resp != nil {
		return
	}

	usecaseErr := controller.updateTankUsecase.Update(ctx, connection, input)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			firstUsecaseError := usecaseErr.PopUsecaseError()
			secondUsecaseError := usecaseErr.PopUsecaseError()

			if secondUsecaseError == nil {
				resp = response.NewControllerError(response.WaterTankInvalidRequest, firstUsecaseError.Error())
				err = firstUsecaseError
				logs.Gateway().Error(err.Error())
				return
			}

			resp = response.NewControllerError(response.WaterTankNotFound, secondUsecaseError.Error())
			err = secondUsecaseError
			logs.Gateway().Error(err.Error())
			return
		default:
			resp = response.NewControllerError(response.WaterTankInternalServerError, usecaseErr.LastUsecaseError().Error())
		}

		err = usecaseErr.LastUsecaseError()
		logs.Gateway().Error(err.Error())
		return
	}

	resp = response.NewControllerEmptyResponse(response.WaterTankNoContent)

	return
}

func (controller *InternalController) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (resp *response.ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank, of group '%s' state...", input.TankName, input.Group))

	resp = validation.Validate(ctx, input, validation.GetTankSchemaLoader)
	if resp != nil {
		return
	}

	usecaseResponse, usecaseErr := controller.getTankUsecase.Get(ctx, connection, input)

	if usecaseErr.HasError() {
		resp = response.SwitchError(usecaseErr)
		return
	}

	resp = response.NewControllerResponse(response.WaterTankOK, usecaseResponse)

	return
}

func (controller *InternalController) GetGroup(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (resp *response.ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank group...", input.Group))

	resp = validation.Validate(ctx, input, validation.GetGroupSchemaLoader)
	if resp != nil {
		return
	}

	usecaseResponse, usecaseErr := controller.getGroupUsecase.Get(ctx, connection, input)

	if usecaseErr.HasError() {
		resp = response.SwitchError(usecaseErr)
		return
	}

	resp = response.NewControllerGroupResponse(response.WaterTankOK, usecaseResponse)

	return
}
