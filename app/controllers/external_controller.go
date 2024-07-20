package controllers

import (
	"fmt"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/get_group"
	"water-tank-api/app/core/usecases/get_tank"
)

type ExternalController struct {
	tank water_tank.WaterTankData
}

func NewExternalController(tank water_tank.WaterTankData) *ExternalController {
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
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
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
			if useCase == get_group.ErrWaterTankMissingGroup {
				response = NewControllerError(WaterTankBadRequest, usecaseErr.LastError().Error())
				return
			}
			response = NewControllerError(WaterTankNotFound, usecaseErr.LastError().Error())
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	response = NewControllerGroupResponse(WaterTankOK, usecaseResponse)

	return
}
