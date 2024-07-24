package controllers

import (
	"fmt"
	"water-tank-api/app/core/entity/logs"
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

func (controller *ExternalController) Get(tank string, group string) (response *ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank, of group '%s' state...", tank, group))

	usecaseResponse, usecaseErr := controller.getTankUsecase.Get(tank, group)

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

	usecaseResponse, usecaseErr := controller.getGroupUsecase.Get(group)

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
