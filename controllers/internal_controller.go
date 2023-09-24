package controllers

import (
	"water-tank-api/core/entity/data"
	register_tank "water-tank-api/core/usecases/register_tank"
	update_tank_state "water-tank-api/core/usecases/update_tank_state"
)

type InternalController struct {
	tank data.WaterTankData
}

func NewInternalController(tank data.WaterTankData) *InternalController {
	return &InternalController{
		tank: tank,
	}
}

func (controller *InternalController) Create(tank string, group string, capacity data.Capacity) (response *ControllerResponse, err error) {
	create := register_tank.NewWaterTank(controller.tank)

	usecaseErr := create.Create(tank, group, capacity)

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

	response = NewControllerResponse(WaterTankNoContent, map[string]interface{}{})

	return
}

func (controller *InternalController) Update(tank string, currentLevel data.Capacity) (response *ControllerResponse, err error) {
	update := update_tank_state.NewWaterTankUpdate(controller.tank)

	usecaseErr, foundErr := update.Update(tank, currentLevel)

	if foundErr != nil {
		response = NewControllerError(WaterTankNotFound, foundErr.Error())
		err = foundErr
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

	response = NewControllerResponse(WaterTankNoContent, map[string]interface{}{})

	return
}
