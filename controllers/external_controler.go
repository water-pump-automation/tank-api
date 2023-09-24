package controllers

import (
	"water-tank-api/core/entity/data"
	"water-tank-api/core/usecases/get_group"
	get_tank "water-tank-api/core/usecases/get_tank"

	"github.com/fatih/structs"
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
			response = NewControllerError(WaterTankNotFound, usecaseErr.UsecaseError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.UsecaseError().Error())
			break
		}

		err = usecaseErr.UsecaseError()
		return
	}

	response = NewControllerResponse(WaterTankOK, structs.Map(usecaseResponse))

	return
}

func (controller *ExternalController) GetAll(group string) (response *ControllerResponse, err error) {
	getUsecase := get_group.NewGetGroupWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(group)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankNotFound, usecaseErr.UsecaseError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.UsecaseError().Error())
			break
		}

		err = usecaseErr.UsecaseError()
		return
	}

	response = NewControllerResponse(WaterTankOK, structs.Map(usecaseResponse))

	return
}
