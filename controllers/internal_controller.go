package controllers

import (
	"water-tank-api/core/entity/data"
)

type InternalController struct {
	tank data.WaterTankData
}

func NewInternalController(tank data.WaterTankData) *InternalController {
	return &InternalController{
		tank: tank,
	}
}

func (controller *InternalController) Create(tank string) (response *ControllerResponse, err error) {
	return
}

func (controller *InternalController) Update(group string) (response *ControllerResponse, err error) {
	return
}
