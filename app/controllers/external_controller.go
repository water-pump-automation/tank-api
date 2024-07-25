package controllers

import (
	"context"
	"fmt"
	"water-tank-api/app/controllers/response"
	"water-tank-api/app/controllers/validation"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/get_group"
	"water-tank-api/app/core/usecases/get_tank"
)

type ExternalController struct {
	getTankUsecase  *get_tank.GetWaterTank
	getGroupUsecase *get_group.GetGroupWaterTank
}

func NewExternalController(
	getTankUsecase *get_tank.GetWaterTank,
	getGroupUsecase *get_group.GetGroupWaterTank,
) *ExternalController {
	return &ExternalController{
		getTankUsecase:  getTankUsecase,
		getGroupUsecase: getGroupUsecase,
	}
}

func (controller *ExternalController) Get(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (resp *response.ControllerResponse, err error) {
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

func (controller *ExternalController) GetGroup(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (resp *response.ControllerResponse, err error) {
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
