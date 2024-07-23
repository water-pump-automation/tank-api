package controllers

import (
	"fmt"
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases"
	register_tank "water-tank-api/app/core/usecases/create_tank"
	update_tank_state "water-tank-api/app/core/usecases/update_tank_state"
)

type Controller struct {
	tank            water_tank.WaterTankData
	externalMethods *ExternalController
}

func NewController(tank water_tank.WaterTankData) *Controller {
	return &Controller{
		tank:            tank,
		externalMethods: NewExternalController(tank),
	}
}

func (controller *Controller) Create(tank string, group string, capacity water_tank.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			tank, group, usecases.ConvertCapacityToLiters(capacity)),
	)

	create := register_tank.NewWaterTank(controller.tank)

	accessToken, usecaseErr := create.Create(tank, group, capacity)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(WaterTankInvalidRequest, usecaseErr.LastError().Error())
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	response = NewControllerCreateResponse(WaterTankOK, accessToken)
	return
}

func (controller *Controller) Update(tank string, group string, accessToken access.AccessToken, currentLevel water_tank.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s', water level to %s",
			tank, group, usecases.ConvertCapacityToLiters(currentLevel)),
	)

	update := update_tank_state.NewWaterTankUpdate(controller.tank)

	usecaseErr := update.Update(tank, group, accessToken, currentLevel)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			firstUsecaseError := usecaseErr.PopError()
			secondUsecaseError := usecaseErr.PopError()

			if secondUsecaseError == nil {
				response = NewControllerError(WaterTankInvalidRequest, firstUsecaseError.Error())
				err = firstUsecaseError
				logs.Gateway().Error(err.Error())
				return
			}

			response = NewControllerError(WaterTankNotFound, secondUsecaseError.Error())
			err = secondUsecaseError
			logs.Gateway().Error(err.Error())
			return
		default:
			response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastError().Error())
		}

		err = usecaseErr.LastError()
		logs.Gateway().Error(err.Error())
		return
	}

	response = NewControllerEmptyResponse(WaterTankNoContent)

	return
}

func (controller *Controller) Get(tank string, group string) (response *ControllerResponse, err error) {
	return controller.externalMethods.Get(tank, group)
}

func (controller *Controller) GetGroup(group string) (response *ControllerResponse, err error) {
	return controller.externalMethods.GetGroup(group)
}
