package controllers

import (
	"fmt"
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/create_tank"
	"water-tank-api/app/core/usecases/get_group"
	"water-tank-api/app/core/usecases/get_tank"
	"water-tank-api/app/core/usecases/ports"
	update_tank_state "water-tank-api/app/core/usecases/update_tank_state"
)

type InternalController struct {
	getTankUsecase    *get_tank.GetWaterTank
	getGroupUsecase   *get_group.GetGroupWaterTank
	createTankUsecase *create_tank.CreateWaterTank
	updateTankUsecase *update_tank_state.UpdateWaterTank
}

func NewInternalController(
	getTankUsecase *get_tank.GetWaterTank,
	getGroupUsecase *get_group.GetGroupWaterTank,
	createTankUsecase *create_tank.CreateWaterTank,
	updateTankUsecase *update_tank_state.UpdateWaterTank,
) *InternalController {
	return &InternalController{
		getTankUsecase:    getTankUsecase,
		getGroupUsecase:   getGroupUsecase,
		createTankUsecase: createTankUsecase,
		updateTankUsecase: updateTankUsecase,
	}
}

func (controller *InternalController) Create(tank string, group string, capacity water_tank.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			tank, group, ports.ConvertCapacityToLiters(capacity)),
	)

	accessToken, usecaseErr := controller.createTankUsecase.Create(tank, group, capacity)

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

func (controller *InternalController) Update(tank string, group string, accessToken access.AccessToken, currentLevel water_tank.Capacity) (response *ControllerResponse, err error) {
	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s', water level to %s",
			tank, group, ports.ConvertCapacityToLiters(currentLevel)),
	)

	usecaseErr := controller.updateTankUsecase.Update(tank, group, accessToken, currentLevel)

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

func (controller *InternalController) Get(tank string, group string) (response *ControllerResponse, err error) {
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

func (controller *InternalController) GetGroup(group string) (response *ControllerResponse, err error) {
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
