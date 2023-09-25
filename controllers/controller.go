package controllers

import (
	"fmt"
	"water-tank-api/core/entity/logs"
	data "water-tank-api/core/entity/water_tank"
	"water-tank-api/core/usecases/get"
	register_tank "water-tank-api/core/usecases/register_tank"
	update_tank_state "water-tank-api/core/usecases/update_tank_state"
)

type Controller struct {
	tank            data.WaterTankData
	externalMethods *ExternalController
}

func NewController(tank data.WaterTankData) *Controller {
	return &Controller{
		tank:            tank,
		externalMethods: NewExternalController(tank),
	}
}

func (controller *Controller) Create(tank string, group string, capacity data.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			tank, group, get.ConvertCapacityToLiters(capacity)),
	)

	create := register_tank.NewWaterTank(controller.tank)

	usecaseErr := create.Create(tank, group, capacity)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankInvalidRequest, usecaseErr.LastError().Error())
			break
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
			break
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	return
}

func (controller *Controller) Update(tank string, currentLevel data.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's water level to %s",
			tank, get.ConvertCapacityToLiters(currentLevel)),
	)

	update := update_tank_state.NewWaterTankUpdate(controller.tank)

	usecaseErr := update.Update(tank, currentLevel)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			firstUsecaseError := usecaseErr.PopError()
			secondUsecaseError := usecaseErr.PopError()

			if secondUsecaseError == nil {
				response = NewControllerError(WaterTankNotFound, firstUsecaseError.Error())
				err = firstUsecaseError
				logs.Gateway().Error(err.Error())
				return
			}

			response = NewControllerError(WaterTankInvalidRequest, secondUsecaseError.Error())
			err = secondUsecaseError
			logs.Gateway().Error(err.Error())
			return
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
			break
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	return
}

func (controller *Controller) Get(tank string) (response *ControllerResponse, err error) {
	return controller.externalMethods.Get(tank)
}

func (controller *Controller) GetAll(group string) (response *ControllerResponse, err error) {
	return controller.externalMethods.GetAll(group)
}
