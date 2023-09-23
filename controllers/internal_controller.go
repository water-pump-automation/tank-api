package controllers

import (
	"water-tank-api/core/entity/data"

	"github.com/fatih/structs"
)

type InternalController struct {
	tank data.WaterTankData
}

func NewInternalController(tank data.WaterTankData) *InternalController {
	return &InternalController{
		tank: tank,
	}
}

func (controller *InternalController) Create(tank string, capacity data.Capacity) (response *ControllerResponse, err error) {
	create := create_tank.NewWaterTank(controller.tank)

	usecaseResponse, usecaseErr := create.Create(tank, capacity)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankBadRequest, usecaseErr.UsecaseError().Error())
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

func (controller *InternalController) Update(tank string, currentLevel data.Capacity) (response *ControllerResponse, err error) {
	update := update_tank.NewWaterTank(controller.tank)

	usecaseResponse, usecaseErr := update.Update(tank, currentLevel)

	if tank == "" {
		response = NewControllerError(WaterTankBadRequest, usecaseErr.UsecaseError().Error())
		err = WaterTankEmptyNameError
		return
	}

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankInvalidRequest, usecaseErr.UsecaseError().Error())
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
