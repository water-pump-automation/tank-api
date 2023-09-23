package controllers

import (
	"water-tank-api/core/entity/data"
	get_tank "water-tank-api/core/usecases/get"
	"water-tank-api/core/usecases/get_group"

	"github.com/fatih/structs"
)

type Controller struct {
	tank data.WaterTankData
}

func NewController(tank data.WaterTankData) *Controller {
	return &Controller{
		tank: tank,
	}
}

func (controller *Controller) Get(tank string) (response *ControllerResponse, err error) {
	getUsecase := get_tank.NewGetWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(tank)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(NetStatNotFound, usecaseErr.UsecaseError().Error())
			break
		default:
			response = NewControllerError(NetStatInternalServerError, usecaseErr.UsecaseError().Error())
			break
		}

		err = usecaseErr.UsecaseError()
		return
	}

	response = NewControllerResponse(NetStatOK, structs.Map(usecaseResponse))

	return
}

func (controller *Controller) GetAll(group string) (response *ControllerResponse, err error) {
	getUsecase := get_group.NewGetGroupWaterTank(controller.tank)

	usecaseResponse, usecaseErr := getUsecase.Get(group)

	if usecaseErr.HasError() {
		switch usecaseErr.EntityError() {
		case nil:
			response = NewControllerError(NetStatNotFound, usecaseErr.UsecaseError().Error())
			break
		default:
			response = NewControllerError(NetStatInternalServerError, usecaseErr.UsecaseError().Error())
			break
		}

		err = usecaseErr.UsecaseError()
		return
	}

	response = NewControllerResponse(NetStatOK, structs.Map(usecaseResponse))

	return
}
