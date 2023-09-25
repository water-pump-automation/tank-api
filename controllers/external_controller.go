package controllers

import (
	data "water-tank-api/core/entity/water_tank"
	get_group "water-tank-api/core/usecases/get/group"
	get_tank "water-tank-api/core/usecases/get/tank"
)

type ExternalController struct {
	tank data.WaterTankData
}

func NewExternalController(tank data.WaterTankData) *ExternalController {
	return &ExternalController{
		tank: tank,
	}
}

func (controller *ExternalController) Get(tank string) (response *ControllerResponse, err error) {
	getUsecase := get_tank.NewGetWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(tank)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankNotFound, usecaseErr.LastError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
			break
		}

		err = usecaseErr.LastError()
		return
	}

	response = NewControllerResponse(WaterTankOK, usecaseResponse)

	return
}

func (controller *ExternalController) GetAll(group string) (response *ControllerResponse, err error) {
	getUsecase := get_group.NewGetGroupWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(group)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankNotFound, usecaseErr.LastError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
			break
		}

		err = usecaseErr.LastError()
		return
	}

	response = NewControllerGroupResponse(WaterTankOK, usecaseResponse)

	return
}
