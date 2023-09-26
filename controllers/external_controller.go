package controllers

import (
	"fmt"
	"water-tank-api/core/entity/logs"
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

func (controller *ExternalController) Get(tank string, group string) (response *ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank, of group '%s' state...", tank, group))

	getUsecase := get_tank.NewGetWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(tank, group)

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
		logs.Gateway().Error(err.Error())
		return
	}

	response = NewControllerResponse(WaterTankOK, usecaseResponse)

	return
}

func (controller *ExternalController) GetGroup(group string) (response *ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank group...", group))

	getUsecase := get_group.NewGetGroupWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(group)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			useCase := usecaseErr.LastError()
			if useCase == get_group.WaterTankMissingGroup {
				response = NewControllerError(WaterTankBadRequest, usecaseErr.LastError().Error())
				return
			}
			response = NewControllerError(WaterTankNotFound, usecaseErr.LastError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
			break
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	response = NewControllerGroupResponse(WaterTankOK, usecaseResponse)

	return
}
